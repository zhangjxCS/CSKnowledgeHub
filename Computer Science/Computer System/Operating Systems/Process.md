## Process

### Introduction

The **process** is the major OS abstraction of a running program. 

- Address space: the memory the process can access
- register: many instructions explicitly read or update registers
  - program counter: which instruction of the program will execute next
  - stack pointer: manage the stack for function parameters
- Persistent disk: OS can access the disk

<img src="https://tva1.sinaimg.cn/large/008vxvgGgy1h7g1k6wmpaj30u00wr40l.jpg" alt="image-20221023194959877" style="zoom:33%;" />

- Running: In the running state, a process is running on a processor.
- Ready: In the ready state, a process is ready to run but for some reason the OS has chosen not to run it at this given moment.
- Blocked: In the blocked state, a process has performed some kind of operation that makes it not ready to run until some other event takes place.

<img src="https://tva1.sinaimg.cn/large/008vxvgGgy1h7g1jwmqhrj30ly0cct9d.jpg" alt="image-20221023194943870" style="zoom:50%;" />

### API

- The **fork()** system call is used in UNIX systems to create a new pro- cess. The creator is called the **parent**; the newly created process is called the **child**.
- The **wait()** system call allows a parent to wait for its child to complete execution.
- The **exec()** family of system calls allows a child to break free from its similarity to its parent and execute an entirely new program.
- Signal: the **signal()** system call catch various signals, ctrl + c sends a SIGINT signal, and ctrl + z sends a SIGTSTP signal.

## Thread

- Multi-threads share the same address space (shared heap, file block)
- Thread has its own thread control blocks(TCB), stack, stack pointer, program counter

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20221009115237311.png" alt="image-20221009115237311" style="zoom:50%;" />

- race condition: the results depend on the timing execution of the code
- critical section: code that accesses a shared variable and must not be concurrently executed by more than one thread
- Indeterminate program: the output of the program varies from run to run, depending on which threads ran

- mutual exclusive: only a single thread ever enters a critical section, thus avoiding races, and resulting in deterministic program outputs

## Synchronization

### Locks

- Mutual exclusion: when the lock work, it prevents multiple threads from entering a critical section
- Fairness: each thread contending for the lock get a fair shot at acquiring it once it is free
- Performance: the time overheads added by using the lock]

```c
// Disable Interrupts Lock
void lock() {
  	DisableInterrupts();
}
void unlock() {
  	EnableInterrupts();
}
```

- allow any calling thread to perform a privileged operation and give too much trust in application
- does not work on multiprocessors
- turning off interrupts can lead to interrupts becoming lost, and lead to serious systems problems

Hardware support: add some hardware support, CPU executes several commands atomically

```c
// test-and-set
int TestAndSet(int *old_prt, int new) {
  	int old = *old_ptr;
  	*old_ptr = new;
  	return old;
}
```

```c
// compare-and-swap
int CompareAndSwap(int *ptr, int expected, int new) {
  	int original = *ptr;
  	if (original == expected)
      	*ptr = new;
  	return original;
}
```

```c
// fetch-and-add
int FetchAndAdd(int *ptr) {
  	int old = *ptr;
  	*ptr = old + 1;
  	return old;
}
```

- Add some hardware support, CPU executes several commands atomically
- Three kinds of atomic operation

```c
// spin lock
void init() {
  	flag = 0;
}
void lock() {
  	while (TestAndSet(&flag, 1) == 1)
      	;
}
void unlock() {
  	flag = 0;
}
```

- the spin lock only allows a single thread to enter the critical section at a time
- spin locks don't provide any fairness guarantees
- Each of thread will spin for the duration of a time slice, a waste of CPU cycles

```c
// yield lock
void init() {
  	flag = 0;
}
void lock() {
  	while (TestAndSet(&flag, 1) == 1)
      	yield();
}
void unlock() {
  	flag = 0;
}
```

- Simply give the CPU when other threads hold the lock to avoid spinning
- Still costly when too many threads waiting for the CPU
- A thread may get caught in an endless yield loop while other threads repeatedly enter and exit the critical section

```c
typedef struct __lock_t {
  	int flag;
  	int guard;
  	queue_t *q;
} lock_t;
void lock_init(lock_t *m) {
  	m->flag = 0;
  	m->guard = 0;
  	queue_init(m->q)
}
void lock(lock_t *m) {
  	while (TestAndSet(&m->guard, 1) == 1)
      	;
  	if (m->flag == 0) {
      	m->flag = 1;
      	m->guard = 0;
    } else {
      	queue_add(m->q, gettid());
      	m->guard = 0;
      	park();
    }
}
```

### Condition Variables

- an explicit queue that threads can put themselves on when some state of execution is not as desired
- wait() call is executed when a thread wishes to put itself to sleep
- Signal() call is executed when a thread has changed something in the program and thus wants to wake a sleeping thread waiting on this condition

```c
int done  = 0;
pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t c  = PTHREAD_COND_INITIALIZER;
void thr_exit() {
    Pthread_mutex_lock(&m);
    done = 1;
    Pthread_cond_signal(&c);
    Pthread_mutex_unlock(&m);
}
void *child(void *arg) {
    printf("child\n");
    thr_exit();
    return NULL;
}
void thr_join() {
    Pthread_mutex_lock(&m);
    while (done == 0)
				Pthread_cond_wait(&c, &m);
		Pthread_mutex_unlock(&m);
}
int main(int argc, char *argv[]) {
    printf("parent: begin\n");
    pthread_t p;
    Pthread_create(&p, NULL, child, NULL);
    thr_join();
    printf("parent: end\n");
    return 0; 
}
```

```
lock(x)
  	while(condition) cv.wait(x)
unlock(x)

lock(x)
    cv.signal()
unlock(x)

wait()
		add to wait queue
		state = blocked
		unlock(x)
		schedule()
		lock()
		
signal()
		if wait queue !empty
				p = remove_queue
				wake_up(p)
```

### Semaphores

- When a task attempts to acquire a semaphore that is unavailable, the semaphore places the task onto a wait queue and puts the task to sleep

- Sem_wait(): decrement the value of semaphore s by one, wait if value of semaphore s is negative
- Sem_post(): increment the value of semaphore s by one, if there are one or more threads waiting, wake one

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20221010172333918.png" alt="image-20221010172333918" style="zoom:50%;" />

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20221010172349525.png" alt="image-20221010172349525" style="zoom:50%;" />

```c
// producer-consumer problem
void *producer(void *arg) {
  	int i;
  	for (i=0; i<loops; i++) {
      	sem_wait(&empty);
      	sem_wait(&mutex);
      	put(i);
      	sem_post(&mutex);
      	sem_post(&full);
    }
}
void *consumer(void *arg) {
  	int i;
  	for (i=0; i<loops; i++) {
      	sem_wait(&full);
      	sem_wait(&mutex);
      	int temp = get();
      	sem_post(&mutex);
      	sem_post(&empty);
      	printf("%d\n", tmp);
    }
}
```

```c
// reader-writer problem
typedef struct _rwlock_t {
  	sem_t lock;      // binary semaphore (basic lock)
  	sem_t writelock; // allow ONE writer/MANY readers
  	int   readers;   // #readers in critical section
  	rwlock_t;
void rwlock_init(rwlock_t *rw) {
  	rw->readers = 0;
  	sem_init(&rw->lock, 0, 1);
  	sem_init(&rw->writelock, 0, 1);
}
void rwlock_acquire_readlock(rwlock_t *rw) {
    sem_wait(&rw->lock);
    rw->readers++;
    if (rw->readers == 1) // first reader gets writelock
      	sem_wait(&rw->writelock);
    sem_post(&rw->lock);
}
void rwlock_release_readlock(rwlock_t *rw) {
    sem_wait(&rw->lock);
    rw->readers--;
    if (rw->readers == 0) // last reader lets it go
      	sem_post(&rw->writelock);
    sem_post(&rw->lock);
}
void rwlock_acquire_writelock(rwlock_t *rw) {
  	sem_wait(&rw->writelock);
}
```

### Deadlock

Deadlock happens whenever circular waiting

1. mutual exclusion
2. hold and wait
3. circular wait
4. no preemption(rollback)
