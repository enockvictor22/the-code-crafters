# Go structs

A struct is basically a custom container used to bundle different data types into one record.

While an array is stuck with just one type, a struct lets me mix things like strings and integers to represent real-world objects, like a "Person" or a "Car."

To set one up, I use the "type" and "struct" keywords to define the blueprint. Once I’ve created a variable from that blueprint, I use the dot operator (.) to assign or grab values from its members. They are also super easy to pass into functions as arguments, which helps keep my code organized when I'm dealing with complex data.

I personally find structs to be the backbone of any serious Go project. They make the code way more readable seeing pers1.name is much more relatable than trying to remember which index in a slice holds a person's name.