# service-skeleton

An opinionated framework for building golang web services tools on top of [`gin-gonic/gin`](https://github.com/gin-gonic/gin) and [`rs/zerolog`](https://github.com/rs/zerolog).

## Why

While [`gin-gonic/gin`](https://github.com/gin-gonic/gin) provides an easy method for engineers to get started with services, there are often missing critical components around creating the server binary, starting the server, proper logging, etc. This can all be cobbled together from various examples, but that is far from ideal for getting started quickly. This project aims to fill that void by implementing a skeleton based upon those provided by the those underlying tools.

### Examples

For examples on how to perform various tasks, see the following examples:

- [`hello-world`](examples/hello-world): The hello-world example.
