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
| is_active   | Boolean       |
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

# Rotas

## [Auth](./docs/AUTH.md)

### `POST /auth/login`

> Acesso: Publico  
> Descrição: Responsável por autenticar o usuário e entregar o token JWT.

### `POST /auth/register`

> Acesso: Admin  
> Descrição: Criar uma nova conta de usuário.

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
