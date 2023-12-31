# golang-with-tests
Learn Golang with TDD Approach

Source: https://quii.gitbook.io/learn-go-with-tests/

## Important Notes

### Defining Variables
We use := to declare and initializing variables which is short hard for both steps.

We use var name type to declare a variable with name and its type without initializing

### Function Name Convention

The function name starts with a lowercase letter.
In Go, public functions start with a capital letter and private ones start with a lowercase.

We don't want the internals of our algorithm to be exposed to the world, so we made this function private.

### Benchmarking

When the benchmark code is executed, it runs b.N times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

### Arrays Initialization

Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:

[N]type{value1, value2, ..., valueN}
numbers := [5]int{1, 2, 3, 4, 5}

[...]type{value1, value2, ..., valueN}
numbers := [...]int{1, 2, 3, 4, 5}

An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.

### Range For Loop

range lets you iterate over an array. On each iteration, range returns two values - the index and the value. We are choosing to ignore the index value by using _ blank identifier.

### Slices

Go has slices which do not encode the size of the collection and instead can have any size.

Go does not let you use equality operators with slices.

You could write a function to iterate over each got and want slice and check their values but for convenience sake, we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.

reflect.DeepEqual is not type safe at compile time!

### Test Coverage

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base.

Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.

Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests.

run go test -cover

### Struct
We can create a simple type using a struct. A struct is just a named collection of fields where you can store data.

### String Formatter
Use of g will print a more precise decimal number in the error message (fmt options). For example, using a radius of 1.5 in a circle area calculation, f would show 7.068583 whereas g would show 7.0685834705770345.

### Function Definition
Some programming languages allow you to have two function with same name but different parameters. But you cant do in go.

We have two choices:
- You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.

- We can define methods on our newly defined types instead.

### Methods
A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

### Interface
Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

Normally you have to write code to say My type Foo implements interface Bar.

In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

### Table Driven Tests
Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.

Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests. They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.


### Pointers

Pointers let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.

In Go, when you call a function or a method the arguments are copied.

When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the address of that bit of memory with &myVal.

Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

### Typealias swift on go

Go lets you create new types from existing ones.
The syntax is type MyName OriginalType

To make Bitcoin you just use the syntax Bitcoin(999).
By doing this we're making a new type and we can declare methods on them.

This can be very useful when you want to add some domain specific functionality on top of existing types.

### nil
nil is synonymous with null from other programming languages. Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

Like null if you try to access a value that is nil it will throw a runtime panic. This is bad! You should make sure that you check for nils.

Pointers can be nil

When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

Useful for when you want to describe a value that could be missing

### t.Fatal

We've introduced t.Fatal which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around. Without this the test would carry on to the next step and panic because of a nil pointer.

### Errors

Errors are the way to signify failure when calling a function/method.

By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.

This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.

Don’t just check errors, handle them gracefully

### Maps

Maps allow you to store items in a manner similar to a dictionary. You can think of the key as the word and the value as the definition.

Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the [].

The key type is special. It can only be a comparable type because without the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value. Comparable types are explained in depth in the language spec.

The value type, on the other hand, can be any type you want. It can even be another map.

Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.
The delete function returns nothing, and we based our Delete method on the same notion.

Since deleting a value that's not there has no effect, unlike our Update and Add methods, we don't need to complicate the API with errors.


### Initialize empty maps

```
var dictionary = map[string]string{}
​
// OR
​
var dictionary = make(map[string]string)
```

### Dependency Injection

The advantages of Injecting a Depencency are:
1. Test our code: If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.

2. Separate our concerns: decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.

3. Allow our code to be re-used in different contexts: The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

### os.Stdout vs bytes.Buffer

We know we want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.

In main we will send to os.Stdout so our users see the countdown printed to the terminal.

In test we will send to bytes.Buffer so our tests can capture what data is being generated.

### Problem facing about having real dependencies on our tests:

1. Every forward-thinking post about software development emphasises the importance of quick feedback loops.

2. Slow tests ruin developer productivity.

3. Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 3s added to the test run for every new test of Countdown?

### Spy
Spies are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc.

### Notes for Mocking:
1. The thing you are testing is having to do too many things (because it has too many dependencies to mock)
-> Break the module apart so it does less

2. Its dependencies are too fine-grained
-> Think about how you can consolidate some of these dependencies into one meaningful module

3. Your test is too concerned with implementation details
-> Favour testing expected behaviour rather than the implementation

### Test Doubles:
The generic term he uses is a Test Double (think stunt double). Test Double is a generic term for any case where you replace a production object for testing purposes. There are various kinds of double that Gerard lists:

1. Dummy objects are passed around but never actually used. Usually they are just used to fill parameter lists.

2. Fake objects actually have working implementations, but usually take some shortcut which makes them not suitable for production (an InMemoryTestDatabase is a good example).

3. Stubs provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.

4. Spies are stubs that also record some information based on how they were called. One form of this might be an email service that records how many messages it was sent.

5. Mocks are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.

### GoRoutines
Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine.

Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.

To tell Go to start a new goroutine we turn a function call into a go statement by putting the keyword go in front of it: go doSomething().

### Data Races:
This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.

### Channels:
We can solve this data race by coordinating our goroutines using channels. Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.