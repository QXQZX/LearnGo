package designpattern

// 享元模式

//在程序设计中，我们常常会碰到一些很重型的对象，它们通常拥有很多的成员属性，
//当系统中充斥着大量的这些对象时，系统的内存将会承受巨大的压力。
//此外，频繁的创建这些对象也极大地消耗了系统的CPU。很多时候，这些重型对象里，
//大部分的成员属性都是固定的，这种场景下， 可以使用享元模式进行优化，
//将其中固定不变的部分设计成共享对象（享元，flyweight），这样就能节省大量的系统内存和CPU。
//享元模式摒弃了在每个对象中保存所有数据的方式， 通过共享多个对象所共有的相同状态，
//让你能在有限的内存容量中载入更多对象。

//当我们决定对一个重型对象采用享元模式进行优化时，首先需要将该重型对象的属性划分为两类，
//能够共享的和不能共享的。
//前者我们称为内部状态（intrinsic state），存储在享元中， 不随享元所处上下文的变化而变化；
//后者称为外部状态（extrinsic state），它的值取决于享元所处的上下文，
//因此不能共享。比如，文章A和文章B都引用了图片A，由于文章A和文章B的文字内容是不一样的，
//因此文字就是外部状态，不能共享；但是它们所引用的图片A是一样的，属于内部状态，
//因此可以将图片A设计为一个享元

//工厂模式通常都会和享元模式结对出现，享元工厂提供了唯一获取享元对象的接口，
//这样Client就感知不到享元是如何共享的，降低了模块的耦合性。享元模式和单例模式有些类似的地方
//，都是在系统中共享对象，但是单例模式更关心的是对象在系统中仅仅创建一次，
//而享元模式更关心的是如何在多个对象中共享相同的状态。
