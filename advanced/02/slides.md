% Building a Web Service
% Tom Arrell
% Tue 6th July, 2021

# Principles

- KISS
- Build it when you need it
- "duplication is far cheaper than the wrong abstraction"

---

# Std lib

- Great support for building web services
- Built in router
- Built in templating library
- Great JSON encoding/decoding support
- Httptest for testing your handlers

---

# Notes

- New goroutine spawned for each request

---

# Hello World

Let's begin by writing a simple Hello World handler, which replies with 200 and
some bytes.

---

# Testing

Let's now see how we can test our handler using the standard library's
`httptest` package.

---

*lesson 2, fin*

**Questions?**
