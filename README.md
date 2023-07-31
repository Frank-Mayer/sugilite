<h1 align="center">
  <img
    width="128px"
    height="128px"
    alt="💎"
    src="https://raw.githubusercontent.com/Frank-Mayer/sugilite/main/favicon.svg"
  />
  <br/>
  <p>sugilite</p>
</h1>

Sugilite is a compiled general purpose programming language.

This project is **work in progress**.

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.

## Memory

Thanks to the automatic memory management, it is easy to write a well-functioning program with Sugilite.

If necessary, the management can be turned off for single variables. These self-managed variables are used as pointers and must be deleted manually.

## Function return values

There are several ways a function can return values.

With the `yank` keyword a. Return value is determined or overwritten. The function is then terminated. This behavior is similar to the `return` from C-like languages.

The `yield` keyword can be used to specify or overwrite a return value of the current function. This does not terminate the function.

## Empty values and errors

To be able to have empty values, the built-in data type `option` is used. `option` can have a value or not.

Errors are mapped with the built-in `result` data type. A `result` can contain a value or an error.

## Credits

[Icon created by max.icons](https://www.flaticon.com/authors/maxicons)
