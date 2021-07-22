#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/time.h>
#include <string.h>

float elapsed_time(struct timeval before, struct timeval after);

struct Node {
	int Data;
	struct Node *Next;
};


struct Node *freeNodeList = NULL;
int malloced_nodes = 0;
int free_list_count = 0;
struct Node *new_node(int value, struct Node *next);
void free_list(struct Node *head);

int preallocated = 0;
struct Node *arrayallocation = NULL;
int preallocated_node_count = 0;
int next_node_index;
void perform_preallocation(int node_count);

struct Node *mergesort(struct Node *head);
struct Node *randomValueList(int n);
int isSorted(struct Node *head);

int
main(int ac, char **av)
{
	int c, n;
	float total = 0.0;
	struct timeval tv;

	gettimeofday(&tv, NULL);
	srandom(tv.tv_usec | getpid());

    while (EOF != (c = getopt(ac, av, "p")))
    {
        switch (c)
        {
        case 'p':
			preallocated = 1;
		}
	}


	for (n = 10000; n < 8000000; n += 200000) {
		int i;

		if (preallocated) {
			perform_preallocation(n);
		}

		for (i = 0; i < 11; i++) {
			struct timeval before, after;

			struct Node *nl, *head = randomValueList(n);

			gettimeofday(&before, NULL);
			nl = mergesort(head);
			gettimeofday(&after, NULL);

			if (i > 0) {
				if (!isSorted(nl)) {
					fprintf(stderr, "n %d, i %d, final list not sorted\n", n, i);
				}
				/* Don't time first sort, warm the cache */
				total += elapsed_time(before, after);
			}

			/*
			printf("n %d, i %d\n", n, i);
			printf("before free list count %d\n", free_list_count);
			printf("Malloced node count %d\n", malloced_nodes);
			fflush(stdout);
			*/
			free_list(nl);
			/*
			printf("after  free list count %d\n", free_list_count);
			printf("Malloced node count %d\n", malloced_nodes);
			fflush(stdout);
			*/
		}

		total /= 10.;
		printf("%d\t%.04f\n", n, total);
		fflush(stdout);
	}
}

int
isSorted(struct Node *head) {
	if (head == NULL || head->Next == NULL) {
		return 1;
	}
	for (; head->Next != NULL; head = head->Next) {
		if (head->Data > head->Next->Data) {
			return 0;
		}
	}
	return 1;
}

struct Node *
randomValueList(int n) {
	struct Node *head = NULL;
	int i;
	for (i = 0; i < n; i++) {
		head = new_node((int)random(), head);
	}
	return head;
}

void append(struct Node *n, struct Node **hd, struct Node **tl) {
	if (!*hd) {
		*hd = n;
		*tl = n;
		return;
	}
	(*tl)->Next = n;
	*tl = n;
}

struct Node *
mergesort(struct Node *head) {
	struct Node *hd, *tl, *p;
	int mergecount, k;

	hd = tl = NULL;
	p = head;
	mergecount = 2; /* just to pass the first for-test */

	for (k = 1; mergecount > 1; k *= 2) {

		mergecount = 0;

		while (p) {
			int i, psize = 0, qsize;
			struct Node *q = p;

			for (i = 0; q && i < k; i++) {
				psize++;
				q = q->Next;
			}

			qsize = psize;

			while (psize > 0 && qsize > 0 && q) {
				if (p->Data < q->Data) {
					append(p, &hd, &tl);
					p = p->Next;
					psize--;
					continue;
				}
				append(q, &hd, &tl);
				q = q->Next;
				qsize--;
			}

			for (; psize > 0 && p; psize--) {
				append(p, &hd, &tl);
				p = p->Next;
			}

			for (; qsize > 0 && q; qsize--) {
				append(q, &hd, &tl);
				q = q->Next;
			}

			p = q;

			mergecount++;
		}

		p = hd;
		head = hd;

		hd = NULL;
		tl->Next = NULL;
		tl = NULL;
	}

	return head;
}

struct Node *
new_node(int value, struct Node *next) {
	if (preallocated) {
		struct Node *p = NULL;
		if (next_node_index < preallocated_node_count) {
			p = &(arrayallocation[next_node_index++]);
			p->Data = value;
			p->Next = next;
		}
		return p;
	}
	struct Node *n;
	if (freeNodeList) {
		n = freeNodeList;
		freeNodeList = freeNodeList->Next;
		--free_list_count;
	} else {
		n = malloc(sizeof(*n));
		malloced_nodes++;
	}
	n->Data = value;
	n->Next = next;
	return n;
}

/* frees entire sorted list at end of benchmarking run */
void
free_list(struct Node *head) {
	if (head == NULL)
		return;
	if (preallocated) {
		memset(arrayallocation, 0, sizeof(struct Node)*preallocated_node_count);
		next_node_index = 0;
		return;
	}
	while (head != NULL) {
		struct Node *tmp = head->Next;
		head->Data = -1;
		head->Next = freeNodeList;
		freeNodeList = head;
		++free_list_count;
		head = tmp;
	}
}

void
really_free_list(struct Node *head) {
	struct Node *tmp;
	if (head == NULL)
		return;
	for (tmp = head; head != NULL; head = tmp->Next) {
		head->Next = freeNodeList;
		freeNodeList = head;
		--free_list_count;
	}
}

/* utility function elapsed_time() */
float
elapsed_time(struct timeval before, struct timeval after)
{
    float r = 0.0;

    if (before.tv_usec > after.tv_usec)
    {
        after.tv_usec += 1000000;
        --after.tv_sec;
    }

    r = (float)(after.tv_sec - before.tv_sec)
        + (1.0E-6)*(float)(after.tv_usec - before.tv_usec);

    return r;
}

void
perform_preallocation(int node_count)
{
	if (!preallocated)
		return;
	if (preallocated_node_count >= node_count)
		return;

	if (arrayallocation) {
		free(arrayallocation);
		arrayallocation = NULL;
	}
	preallocated_node_count = node_count;
	next_node_index = 0;

	arrayallocation = calloc(preallocated_node_count, sizeof(struct Node));
}
