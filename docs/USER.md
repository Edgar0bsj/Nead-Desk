## Rota: User

### `GET /admin/users`

- **Descrição:** Gerenciamento e consulta dos usuários.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

- `Body: Vazio`

---

### Response happy

#### `200 Ok` - Sucesso

```json
[
  {
    "id": "5b3d6464-01ea-429e-9a48-baa7c8cf6261",
    "name": "João Silva",
    "email": "joao@email.com",
    "role": "user",
    "created_at": "2026-08-20T10:00:00-03:00"
  }
]
```

---

### `GET /admin/users/:id`

- **Descrição:** Consultar os dados de um usuário específico.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

- `Body: Vazio`

---

### Response happy

#### `200 Ok` - Sucesso

```json
{
  "id": "user-001",
  "name": "João Silva",
  "email": "joao@email.com",
  "role": "user",
  "created_at": "2026-08-20T10:00:00-03:00",
  "updated_at": "2026-08-25T10:00:00-03:00"
}
```

---

### `PATCH /admin/users/:id`

- **Descrição:** Alterar informações administrativas de um usuário.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

```json
{
  "name": "João da Silva",
  "role": "admin"
}
```

---

### Response happy

#### `202 Accepted` - Sucesso

```json
{
  "id": "user-001",
  "name": "João da Silva",
  "email": "joao@email.com",
  "role": "admin",
  "updated_at": "2026-08-25T21:20:00-03:00"
}
```
