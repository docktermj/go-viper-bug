# go-viper-bug

To recreate the bug:

1. Clone
   Example:

    ```console
    git clone git@github.com:docktermj/go-viper-bug.git
    ```

1. Successful test.
   Example:

    ```console
    $ go run main.go cmd3 --database-url postgresql://example.com

    cmd3 database-url: postgresql://example.com
    ```

1. Failed test.
   Example:

    ```console
    $ go run main.go cmd2 --database-url postgresql://example.com

    cmd2 database-url:
    ```

   Note that when the `cmd2` subcommand is used, the value of `--database-url` is not recognized by Viper.
