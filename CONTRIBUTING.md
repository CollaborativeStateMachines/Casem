# Contributing to the Project

Thank you for your interest in contributing! To maintain a high-quality codebase, we follow a strict set of technical and procedural guidelines. Please ensure your contributions align with the following standards.

---

## 🛠 Engineering Standards

### C++ Style Guidelines
We adhere to a "good taste" philosophy. Beyond typical clean code practices, we enforce the following:

* **RAII & Const-Correctness:** These are non-negotiable. Manage resources through object lifetime and ensure logical constancy across the API.
* **Trailing Return Types:** All function declarations must use trailing return type syntax:
    ```cpp
    auto calculate_metrics(int value) -> double;
    ```
* **Exception Safety:** Mark all non-exception-throwing code paths as `noexcept`.
* **Early Escapes:** Avoid deeply nested branches. Prefer `if (condition) return;` to keep the primary logic at a shallow indentation level.
* **Defensive Programming:** Write code that anticipates and handles invalid states or inputs gracefully.
* **Consistency:** Above all, match the surrounding code. New contributions should seat naturally into the existing architecture.

### Tooling
* **Formatting:** We use **Clang Format**. Before submitting, ensure your code matches the project's `.clang-format` specification.
* **Originality:** We value the craft of programming. While Generative AI is a valid productivity tool, we expect you to understand and own the logic you submit.

---

## 🧪 Testing & Quality
* **Test Coverage:** Every pull request must include comprehensive tests.
* **Pre-flight Check:** All tests must pass locally before a PR is submitted. PRs with failing tests or decreased coverage will not be considered for review.

---

## 📝 Commits & Pull Requests

We use **Conventional Commits** for automated changelog generation and history clarity.

We write commit messages in present participle, and full lower-case.

### Commit Messages
Every commit must follow the format: `<type>[optional scope]: <description>`
* *Example:* `feat(parser): adding support for inclusion via urls`
* *Example:* `fix(core): resolving memory leak in resource manager`

### Pull Request Process
1.  **PR Title:** Must also follow **Conventional Commits**.
2.  **Merging:** We use a **rebase-and-merge**. Because the PR title becomes the final commit message on the main branch, ensure it is descriptive and correctly formatted.
3.  **Review:** Once tests pass and style requirements are met, maintainers will review the logic and taste of the implementation.

---

*If you find that the rigor of these standards makes programming less enjoyable, this may not be the right environment for your contributions. We build for performance, safety, and longevity.*