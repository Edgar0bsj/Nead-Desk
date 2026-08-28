## <img src="docs/NeadDeskLogo.png" alt="Exemplo imagem">

## Entidade: User

| Campos        | Type          |
| ------------- | ------------- |
| id            | UUID / String |
| name          | String        |
| email         | String        |
| password_hash | String        |
| role          | UserRole      |
| created_at    | Time          |
| updated_at    | Time          |

### Type UserRole

| CONST     | String  |
| --------- | ------- |
| RoleUser  | "user"  |
| RoleAdmin | "admin" |

---

## Entidade: Tickets

| Campos      | Type           |
| ----------- | -------------- |
| id          | UUID / String  |
| user_id     | UUID / String  |
| assigned_to | UUID / String  |
| category_id | UUID / String  |
| title       | String         |
| description | String         |
| status      | TicketStatus   |
| priority    | TicketPriority |
| created_at  | Time           |
| updated_at  | Time           |
| closed_at   | Time           |

### Type TicketStatus

| CONST            | String        |
| ---------------- | ------------- |
| StatusOpen       | "open"        |
| StatusInProgress | "in_progress" |
| StatusResolved   | "resolved"    |
| StatusClosed     | "closed"      |

### Type TicketPriority

| CONST            | String     |
| ---------------- | ---------- |
| PriorityLow      | "low"      |
| PriorityMedium   | "medium"   |
| PriorityHigh     | "high"     |
| PriorityCritical | "critical" |

---

## Entidade: Categories

| Campos      | Type          |
| ----------- | ------------- |
| id          | UUID / String |
| name        | String        |
| description | String        |
| created_at  | Time          |
| updated_at  | Time          |

---

## Entidade: ticket_comments

| Campos     | Type          |
| ---------- | ------------- |
| id         | UUID / String |
| ticket_id  | UUID / String |
| user_id    | UUID / String |
| content    | String        |
| created_at | Time          |

---

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

---

## Autor

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Edgar0bsj" title="Autor">
        <img src="https://avatars.githubusercontent.com/u/180589510?v=4" width="100px;" alt="Foto do EdgarJr no GitHub"/><br>
        <sub>
          <b>Edgar Junior</b>
        </sub>
      </a>
    </td>

  </tr>
</table>
