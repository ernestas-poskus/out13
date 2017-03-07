+++
tags = [
  "nginx",
]
date = "2017-03-08T08:44:03+02:00"
title = "Serving fresh assets using Nginx location rewrite"
Description = "Serving fresh hashed assets using Nginx location blocks with rewrite."

+++

Recently I have stumbled upon a problem to serve fresh/new assets for user web application.

As Phil Karlton said:

> There are only two hard things in Computer Science: cache invalidation and naming things.

Historically fresh assets problem was approached either by appending appending
url query params (?v=20130102) or renaming/hashing asset file completely (/css/default-2j9alkjan2k2.css).

Former is most popular one but not elegant since it brings explicit dependency
for backend application what fresh/new asset file to include thus requires exact
name file to be present on web server.

This draws 5 main disadvantages of completely hashed asset name:

 1. explicit dependency to include
 2. no fallback mechanism
 3. hashed asset name
 4. exact file name presence on web server
 5. removal of stale assets

I came with solution to use Nginx rewrite block that implicitly drops hash of requested file and serves requested asset.

```nginx
   location @css_assets {
       rewrite ^/css/(.*)\..*\.(.*)$ /css/$1.$2 last;
   }
   location /css/ {
       try_files $uri $uri/ @css_assets;
   }
```

First `location /css/ {` block matches path of /css/ which later executes
`try_files` followed by `location @css_assets {` location block.

Secondly this rewrite `rewrite ^/css/(.*)\..*\.(.*)$ /css/$1.$2 last;` matches
beginning of /css/ path followed by 2 tracked matches.

> /css/app.117c7f2fa4b6ea7a2c077a3dbc9662e6b1c278bd.css

In above example first match tracks (app) and second one (css).
Matched information constructs new implicitly requested file like below.

> /css/app.css

Newly constructed file will be processed by Nginx without redirects and browser
knowing original file name.

### Constructing asset hash

To tell your application which asset must be served use ENVIRONMENT variable and checksum
of asset to be included. Or you can dynamically invalidate/create asset hash for
example hourly or daily depending on release cycle.

> /css/app.$CSS_ASSET_HASH.css

In my case I use simple function below.

```rust
lazy_static! {
    static ref CSS_ASSETS_HASH: String = {
        match env::var("CSS_ASSETS_HASH") {
            Ok(hash) => format!(".{}.", hash),
            Err(_) => ".".to_string(),
        }
    };
}

html! {
  (Css(format!("/css/app{}css", *CSS_ASSETS_HASH)))
}
```

Above example is actual code used in production. It tries to statically initialize
`CSS_ASSETS_HASH` variable, if expected environment is not defined it fallbacks to
dot `.` else it appends 2 dots between supplied environment variable `.ENVIRONMENT_VARIABLE.`.

My solution eliminates almost all of main disadvantages of most popular way of asset
inclusion.

 1. ~~explicit dependency to include~~ - depends if hash is generated
    dynamically or real checksum of asset file is used.
 2. ~~no fallback mechanism~~ - if environment is not defined or any hash is
    supplied it still fallbacks to original requested asset due to regex
    catchall.
 3. ~~hashed asset name~~ - asset name is explicit and easily understood
 4. exact file name presence on web server
 5. ~~removal of stale assets~~ - only original file is deployed


Regards.
