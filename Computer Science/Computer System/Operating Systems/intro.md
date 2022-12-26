## Concepts

- User: applications running with low privileges
- Kernel: the part of the operating system that runs with higher privileges
- User mode: code that runs in user mode has certain limitations
- Kernel mode: code that runs in kernel mode can fully [[1\]](https://linux-kernel-labs.github.io/refs/heads/master/lectures/intro.html#hypervisor) control the CPU
- User space: the memory area reserved to a particular user process
- Kernel space: the memory area that is reserved to the kernel

## Architecture

### Kernel Type

![../_images/ditaa-48374873962ca32ada36c14ab9a83b60f112a1e0.png](https://linux-kernel-labs.github.io/refs/heads/master/_images/ditaa-48374873962ca32ada36c14ab9a83b60f112a1e0.png)

- Monolithic kernel:  there is no access protection between the various kernel subsystems and where public functions can be directly called between various subsystems
- Micro kernel: large parts of the kernel are protected from each-other, usually running as services in user space.

### Address Space

- physical address space: the RAM and device memories are visible on the memory bus
- virtual address space: the CPU sees the memory when the virtual memory module is activated

![../_images/ditaa-a5f93e0d17ccdc2ba24828b620d7227f7fc75e33.png](https://linux-kernel-labs.github.io/refs/heads/master/_images/ditaa-a5f93e0d17ccdc2ba24828b620d7227f7fc75e33.png)

virtual address space is shared between user processes and the kernel, the kernel creates mappings that prevent access to the kernel space from user mode.

