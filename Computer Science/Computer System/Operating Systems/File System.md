## Introduction

- superblock: information about the filesystem instance such as the block size, the root inode, filesystem size
- file: information about an opened file such as the current file pointer
- inode: identifies a file in a unique way and has various properties such as the file size, access rights, file type, etc
- dentry: associates a name with an inode

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20221202104651553.png" alt="image-20221202104651553" style="zoom:50%;" />

- the superblock contains information about the block size as well as the IMAP, DMAP, IZONE and DZONE areas.
- the IMAP area is comprised of multiple blocks which contains a bitmap for inode allocation; it maintains the allocated/free state for all inodes in the IZONE area
- the DMAP area is comprised of multiple blocks which contains a bitmap for data blocks; it maintains the allocated/free state for all blocks the DZONE area

<img src="https://tva1.sinaimg.cn/large/008vxvgGgy1h8pwg2vgmpj317o0f0wff.jpg" alt="image-20221202104902495" style="zoom:50%;" />

## File Operation

Virtual File System: in order to support multiple filesystem types and instances

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20221202105835152.png" alt="image-20221202105835152" style="zoom:50%;" />

### Mount

- input: a storage device
- output: dentry pointing to the root directory
- steps
  - check device
  - determine filesystem parameters
  - locate the root inode

### Open Files

- input: path
- output: file descriptor
- steps:
  - determine the filesystem type
  - for each entry in the path, lookup parent dentry, load inode, load data, find dentry
  - Create a new *file* that points to the last dentry
  - Find a free entry in the file descriptor table and set it to file

### Query Attributes

- input: path
- output: file attributes
- steps:
  - access file, dentry, inode
  - reda file attributes from inode

### Read Data

- Input: file descriptor, offset, length
- output: data
- steps:
  - Access file, dentry, inode
  - determine data blocks
  - copy data blocks to memory

### Write Data

- input: file descriptor, offset, length, data
- steps:
  - Allocate one or more data blocks
  - Add the allocated blocks to the inode and update file size
  - Copy data from userspace to internal buffers and write them to storage

### Close File

- input: file descriptor
- steps:
  - set the file descriptor entry to NULL
  - Decrement file reference counter
  - When the counter reaches 0 free file

### Create File

- input: path
- steps:
  - Determine the inode directory
  - Read data blocks and find space for a new dentry
  - Write back the modified inode directory data blocks

### Delete File

- input: path
- steps:
  - determine the parent inode
  - read parent inode data blocks
  - find and erase the dentry (check for links)
  - when last file is closed: deallocate data and inode blocks