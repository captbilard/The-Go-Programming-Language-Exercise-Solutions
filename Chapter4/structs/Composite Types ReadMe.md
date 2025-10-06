## Struct Embedding & Anonymous Fields
Go gives us the opportunity to use an anonymous field of another struct type in
our struct, what this does is that it uses the law of composition to promote the
fields of the anonymous struct type up to the current struct type. Making it convenient
for us to do things like `x.f` rather than `x.d.e.f`.

In more plain english, struct embedding: Embedding lets you "include" one struct 
inside another without giving it a field name. 
The embedded struct's fields and methods get "promoted" to the outer struct.

```
type Engine struct {
    Horsepower int
    Type       string
}

func (e Engine) Start() {
    fmt.Println("Engine starting...")
}

type Car struct {
    Engine  // embedded - no field name!
    Brand   string
}

func main() {
    c := Car{
        Engine: Engine{Horsepower: 300, Type: "V8"},
        Brand:  "Ford",
    }
    
    // Can access embedded fields directly
    fmt.Println(c.Horsepower)  // 300 (promoted from Engine)
    fmt.Println(c.Type)        // "V8" (promoted from Engine)
    
    // Can also access via the type name
    fmt.Println(c.Engine.Horsepower)  // 300
    
    // Methods are promoted too!
    c.Start()  // "Engine starting..."
}

```

Key points:

- It's not inheritance - it's composition
- The outer struct doesn't "become" the inner type
- If there's a naming conflict, the outer struct's field wins
- You can embed interfaces too (common pattern)

## Composition vs Inheritance
Inheritance follows the "is-a" relationship, while Composition follows an "has-a" relationship
In inheritance, a child inherits everything from the parent whether they need it or not.
In the example below Penguin inherits the `fly()` method even though it can't fly.
```
Animal
├── Bird
│   ├── Penguin  // Can fly()? No! 😱
│   └── Sparrow
└── Fish
    └── FlyingFish  // Can fly()? Yes! 😱😱
```

Composition
Object contains  other objects as parts based on what they need. In the example
below, each animal gets exactly what it needs, nothing more.
```
type Swimmer interface {
    Swim()
}

type Flyer interface {
    Fly()
}

type Penguin struct {
    swimAbility Swimmer  // Has swimming
    // No flying ability - doesn't need it!
}

type Duck struct {
    swimAbility Swimmer  // Has both!
    flyAbility  Flyer
}

type Sparrow struct {
    flyAbility Flyer  // Only flying
}
```

#### Why Go prefers composition:

- More flexible
- Easier to reason about
- No confusing inheritance chains
- Favor explicit over implicit

"Prefer composition over inheritance" is a well-known design principle

#### Real-World Analogy
Inheritance: You're born into a family. You get your parents' traits whether
you want them or not. Can't change your biological parents.

Composition: You build a computer. You choose the CPU, RAM, GPU you want.
Don't like the GPU? Swap it out. Each part is independent.

#### When to Use What?
Inheritance best used when:

- There is a true "is-a" relationships that won't change
- Should be shallow hierarchies (1-2 levels max)

Example: Circle is-a Shape (geometry is well-defined)

Composition (preferred):

"Has-a" or "uses-a" relationships
- Need flexibility
- Building complex systems from simpler parts
- Pretty much everywhere else!

The Mantra: "Favor composition over inheritance"