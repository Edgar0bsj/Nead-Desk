## Rota: Categories

### `POST /admin/categories`

- **Descrição:** Criação de categorias.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

```json
{
  "name": "Rede",
  "description": "Problemas relacionados à rede."
}
```

---

### Response happy

#### `201 Created` - Sucesso

```json
{
  "id": "category-003",
  "name": "Rede",
  "description": "Problemas relacionados à rede.",
  "is_active": true,
  "created_at": "2026-08-25T21:30:00-03:00"
}
```

---

### `GET /admin/categories`

- **Descrição:** Listar categorias.
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
    {
      "id": "377aa3e8-79d5-4402-9086-a1dccaee78ff",
      "name": "Acesso",
      "description": "Problemas relacionados a login.",
      "is_active": true
    },
    {
      "id": "ad2bec2e-05e3-4f49-b8b8-89c4bb74e051",
      "name": "Hardware",
      "description": "Problemas relacionados aos equipamentos.",
      "is_active": true
    }
}
```

---

### `PATCH /admin/categories/:id`

- **Descrição:** Editar categoria.
- **Acesso:** `ADMIN`.

---

### Requisição

**Headers Necessários:**

- `Content-Type: application/json`
- `Authorization: Token Bearer`

**Corpo da Requisição (JSON):**

```json
{
  "name": "Infraestrutura",
  "description": "Problemas de infraestrutura."
}
```

---

### Response happy

#### `200 Ok` - Sucesso

```json
{
  "id": "ad2bec2e-05e3-4f49-b8b8-89c4bb74e051",
  "name": "Infraestrutura",
  "description": "Problemas de infraestrutura.",
  "is_active": true,
  "updated_at": "2026-08-25T21:35:00-03:00"
}
```
