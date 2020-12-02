# Week02

## Q1:
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## A:
不用，感觉可以更加灵活, sql.ErrNoRows返回空值即可， 这里的error只是一个value。
如果是数据库连接出了问题，则依据视屏中所说，wrap error返回堆栈信息。
