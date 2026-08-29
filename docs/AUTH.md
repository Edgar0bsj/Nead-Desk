## Rota: Autenticação

### `POST /auth/login`

**Descrição:** Responsável por autenticar o usuário e entregar o token JWT.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`

**Corpo da Requisição (JSON):**

```json
{
  "email": "joao@email.com",
  "password": "123456"
}
```

---

### Responses

#### `201 Created` - Sucesso

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "token_type": "Bearer"
}
```

#### `401 Unauthorized` - Falha de Autenticação

```json
{
  "error": "invalid credentials"
}
```

#### `500 Internal Server Error` - Erro no Servidor

```json
{
  "status": "error",
  "code": "INTERNAL_ERROR",
  "message": "Ocorreu um erro inesperado no servidor."
}
```

---

## Rota: Usuários criação

### `POST /auth/register`

- **Descrição:** Criar uma nova conta de usuário.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

```json
{
  "name": "João Silva",
  "email": "joao@email.com",
  "password": "123456"
}
```

---

### Response

#### `201 Created` - Sucesso

```json
{
  "id": "12345678-abcd-4567-8901-123456789abc",
  "name": "João Silva",
  "email": "joao@email.com",
  "role": "user",
  "created_at": "2026-08-25T20:00:00-03:00"
}
```
