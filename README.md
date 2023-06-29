# Talks

## Running the slideshows online

Simply click on the link of interest below.

Disclaimer: some images may not load.

### 2023

- [Useful functional-options tricks for better libraries (GopherCon EU 2023)](https://talks.godoc.org/github.com/jub0bs/talks/2023/06/functional-options/main.slide)

### 2022

- [Introduction to Go generics (Go SXB meetup, Dec. 2022)](https://talks.godoc.org/github.com/jub0bs/talks/2022/12/intro-to-generics/main.slide)

## Running the talks locally

You can run the slideshows on your machine after following a few simple steps:

1. If you haven't already installed Go on your machine, do so by
    following the official [installation instructions][install].

2. Important: make sure that the directory where Go installs binaries
    is in your `PATH` environment variable.
    The directory in question is usually given by the following command:

    ```shell
    echo `go env GOPATH`/bin
    ```

3. Install the [`present`][present] tool:

    ```shell
    go install golang.org/x/tools/cmd/present@latest
    ```

4. Clone this repository
  and `cd` into the folder containing the slideshow of interest, e.g.

    ```shell
    git clone https://github.com/jub0bs/talks
    cd talks/2023/06/functional-options
    ```

5. Run the present tool:

    ```shell
    present
    ```

    The output should contain some local URL (like `http://127.0.0.1:3999`).

6. Visit the URL from step 5 in your browser,
    then click on `main.slide` to start the slideshow.

[install]: https://go.dev/doc/install
[present]: https://pkg.go.dev/golang.org/x/tools/cmd/present
