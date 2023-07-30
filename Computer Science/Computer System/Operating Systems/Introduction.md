## Introduction

Operating system is a piece of software that abstracts and arbitrates the underlying hardware system.

- Hide harware complexity
- Manage underlying hardware resources
- Provide isolation and protection

### OS Elements

- Abstraction: process, thread, file, socket, memory page
- Mechanism: create, schedule, open, write, allocate
- Policy: LRU, LFU, Round-Robin

## OS Protection

### OS Mode

- User: applications running with low privileges
- Kernel: the part of the operating system that runs with higher privileges

### User Kernel Interactions

- Trap: when application runs privilege intruction, the application will be trapped, and OS will take control
- System call: application may call system call functions to do privilege instruction
- Signal: OS sends notification to the application

## Architecture

### Kernel Type

- Monolithic kernel: everything included in the OS
- Modular kernel: modular kernel has a basic set of services and APIs that come with it. Anything not included can be added, as a module. As a result, each application can interface with the operating systems in the ways that make most sense to it.
- Micro kernel:  only require the most basic operating system components. Everything else (including file systems and disk drivers) will run outside of the operating system at user-level.

### Linux Architecture

- Hardware
- Linux Kernel
- Standard Libraries (sys call interface)
- Utility program (shells/compiler)
- User applications

