# Concurrent Library Management System

## Features
- Book borrowing
- Book reservation
- Concurrent reservation handling
- Auto-cancel reservations after 5 seconds

## Concurrency Tools Used

### Goroutines
Used to process multiple reservation requests simultaneously.

### Channels
Reservation requests are sent through a channel queue.

### Mutex
sync.Mutex prevents race conditions when updating book data.

### Timer
Reservations expire automatically after 5 seconds.

## Flow

User Request
      ↓
Reservation Channel
      ↓
Worker Goroutine
      ↓
Reserve Book
      ↓
5s Timer
      ↓
Auto Cancel if not borrowed