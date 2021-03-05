require'strict'.on()

box.cfg{
    listen = 3000
}

box.schema.user.grant('guest', 'read,write,execute', 'universe', nil,{
    if_not_exists=true
})

--Создание спейса юзера
user = box.schema.create_space('user', {
    if_not_exists = true,
})

user:format(
        {
            {name = 'user_name', type = 'string'},
            {name = 'email', type = 'string'},
            {name = 'password', type = 'string'},
            {name = 'executor', type = 'boolean'},
            {name = 'description', type = 'string', is_nullable=true},
            {name = 'specializes', type = 'array', is_nullable=true},
            {name = 'img_url', type = 'string', is_nullable = true},
        }
)

user:create_index('primary', {
    type = 'HASH',
    parts = {'user_name'},
    if_not_exists = true,

})
user:create_index('email_index', {
    type = 'TREE',
    unique = true,
    parts = {'email'},
    if_not_exists = true,
})

--Создание спейса заказов
order = box.schema.create_space('order', {
    if_not_exists = true
})

order:format(
        {
            {name = 'order_name', type = 'string'},
            {name = 'customer_name', type = 'string'},
            {name = 'description', type = 'string'},
            {name = 'specializes', type = 'array' },
        }
)

order:create_index('primary', {
    type = 'HASH',
    parts = {'order_name'},
    if_not_exists = true,
})
order:create_index('customer_index', {
    type = 'TREE',
    parts = {'customer_name'},
    if_not_exists = true,
})



--Создание спейса специализаций
specialize = box.schema.create_space('specialize', {
    if_not_exists = true
})

specialize:format(
        {
            {name = 'specialize_name', type = 'string'},
            {name = 'customer_list', type = 'array'},
            {name = 'order_list', type = 'array'},
        }
)

order:create_index('primary', {
    type = 'HASH',
    parts = {'specialize_name'},
    if_not_exists = true,
})

require'console'.start()