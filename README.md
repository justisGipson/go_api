# Project Bluebird



This is to serve as a replacement API for the [`curriculumApp`](https://github.com/CodeliciousProduct/curriculumApp)

No more Firebase/Firestore

Moving to a containerized architecture on GCP

This was written in [Golang](https://go.dev/) using [Fiber](https://github.com/gofiber/fiber) which is an Express inspired framework for the web, also written in Go

Why Go? Because it was built for the cloud. Tons of cloud services run on Go and work well with a Go codebase. And its fast, compiles quickly, type checked, garbage collected.

Even if you're not familiar with golang, this shouldn't be too difficult to get into. Very similar to an express app.

<!-- TODO: compose all cmds for starting the app/containers -->

Bluebird to-dos:
- [x] lesson model
- [ ] course model
- [ ] curriculum model
- [ ] resources model
- [ ] standards model
- [ ] users model
- [x] lesson controller
- [ ] course controller
- [ ] curriculum controller
- [ ] resources controller
- [ ] standards controller
- [ ] users controller
- [x] lesson queries
- [ ] course queries
- [ ] curriculum queries
- [ ] resources queries
- [ ] standards queries
- [ ] users queries
- [x] up migrations for lessons
- [ ] up migrations for curriculum
- [ ] up migrations for courses
- [ ] up migrations for resources
- [ ] up migrations for standards
- [ ] up migrations for users
- [x] down migrations for courses
- [x] down migrations for lessons
- [x] down migrations for curriculum
- [x] down migrations for resources
- [x] down migrations for standards
- [ ] down migrations for users
- [ ] dockerfiles
- [x] Makefile
- [x] private routes
- [ ] public routes
- [x] db connections
- [ ] filters
- [x] swagger docs
- [x] swagger route
- [ ] update documentation (README - ongoing)
- [ ] go docstrings where it is helpful
