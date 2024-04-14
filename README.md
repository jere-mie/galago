# Galago

[Galago](https://en.wikipedia.org/wiki/Galago) is a simple [static site generator](https://en.wikipedia.org/wiki/Static_site_generator) written in [Go](https://go.dev/), with [Jinja2](https://palletsprojects.com/p/jinja/)-like templating.

Templating is handled by the [pongo2](https://github.com/flosch/pongo2) library. While pongo2 attempts to be as similar to Django/Jinja2 as possible, it may not be perfect. Check out its [GitHub repository](https://github.com/flosch/pongo2) for more info on what is and isn't supported.

## Who is This For?

Galago was built for people who want something in between writing bare html and a more opinionated SSG like Hugo. 

### Powerful

Tired of copying and pasting the same navbar code or HTML boilerplate for each page in your site? You can easily use features like extending templates, including files, and macros to help you streamline your development process. 

### Simple and Unopinionated

Forget searching for the right theme for your site or spending hours trying to figure out how to build one from scratch - Galago *just works* with however you want to build your site. No more figuring out what the difference is between `baseof.html`, `single.html`, and `list.html`

---

If you've ever wished you could use the same templating features from Django or Flask for your static sites, then Galago is for you.

## How Does it Work?

Galago will search your `./pages` directory for templates, and render each into its own page, outputting it to the `./public` directory.

It also copies everything in `./static` into the `./public/static` directory.

## Using Base Templates and Macros

If you want to create base templates, macros, etc. that are used in pages but are not rendered as pages themselves, you should create them in a folder other than the `./pages` directory. 

For organization's sake, you can use the `./templates` directory for this. However, Galago does not care where you put them, and they can be placed anywhere.

For example, you may have a `./pages/index.html` page that extends the `./templates/base.html` template and uses the `./templates/macros/user_details.html` macro.

## Downloading

You can find pre-built Galago binaries for Linux and Windows on the Galago repo's [releases page](https://github.com/jere-mie/galago/releases/latest) From there, you can download the binaries and add them to your system's PATH variable.

If you prefer downloading via the cli, you can use the following command to download the latest Galago binary on **Windows**:

```sh
irm -Uri https://github.com/jere-mie/galago/releases/latest/download/galago.exe -O galago.exe
```

and the following command on **Linux**:

```sh
curl -LO https://github.com/jere-mie/galago/releases/latest/download/galago
```

Galago is a lightweight, static binary (under 10Mb), so if you prefer you can install it in your project's directory and run it with `./galago` instead of adding it to your PATH.
