# Food advisor

App to generate menus using ChatGPT

## Requirements

- Docker and docker compose
- ChatGPT API token. See [API readme](./back/README.md) to see how to configure

## Run dev

```bash
docker compose up -d
```

Got to [http://localhost:5173](http://localhost:5173)

## ChatGPT

Check API usage on https://platform.openai.com/usage

## Task list

- API endpoint to get menus and shopping list
- Authentication
- Save selected menus (local storage? database?)
- Relaunch query to replace some menu
- profile where we can define allergy, veggie, ...
