## geek/go-err

`我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？`

`DAO层：dao层叫数据访问，全称为data access object，属于一种比较底层，比较基础的操作，具体到对于某个表、某个实体的增删改查。DAO层一定是和数据库的某几张表一一对应的，其中封装了增删改查基本操作，只做原子操作，增删改查。
`
