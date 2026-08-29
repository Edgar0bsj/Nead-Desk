## Rota: Tickets

### `POST /tickets`

- **Descrição:** Permitir que um usuário registre um problema ou solicitação.
- **Acesso:** `ADMIN ou USER`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

```json
{
  "title": "Não consigo acessar o sistema",
  "description": "Ao tentar acessar aparece usuário ou senha inválidos.",
  "category_id": "8f6c3b6e-1a7d-4c8a-9e31-123456789abc",
  "priority": "high"
}
```

---

### Response happy

#### `201 Created` - Sucesso

```json
{
  "id": "b8d7f9a2-4e31-4c7b-91f2-123456789abc",
  "title": "Não consigo acessar o sistema",
  "description": "Ao tentar acessar aparece usuário ou senha inválidos.",
  "status": "open",
  "priority": "high",
  "category": {
    "id": "8f6c3b6e-1a7d-4c8a-9e31-123456789abc",
    "name": "Acesso"
  },
  "created_at": "2026-08-25T20:10:00-03:00",
  "updated_at": "2026-08-25T20:10:00-03:00"
}
```
