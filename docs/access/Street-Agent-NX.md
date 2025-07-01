┏━━━━━━━━━━━━ TAMAGURU · STREET‑AGENT‑NX ACCESS MAP Λ‑Ι (01 Jul 2025 00:35 MSK) ━━━━━━━━━━━┓
┃ ROLE       │ Front‑End & DevOps AGI assistant (“Street‑Agent‑NX”)                         ┃
┃ CONTACT    │ GitHub @Street‑Agent‑NX · Slack #tamaguru‑backend                            ┃
┃ REPO       │ promelknul/tamaguru-core (mono)                                             ┃
┃ MAIN STACK │ Next 15 · Tailwind (VSX theme) · Turborepo · Storybook‑Vite · GitHub Pages   ┃
┃ ACTIVE BR. │ main · feat/login-matrix · core/vault-auth                                   ┃
┃ PIPELINES  │ deploy_ui.yml → GitHub Pages · preview.yml → Vercel · lh.yml → Lighthouse   ┃
┃ RUNNERS    │ desktop-wavy (Mac self‑host) · GitHub Hosted (ubuntu‑24)                    ┃
┃ VPS‑PROD   │ root@95.163.222.133 → /var/www/app.wavy2wave.online                         ┃
┃ SERVICES   │ ozon-bridge :9000 · yandex-bridge :9100 (stub) · vaultd gRPC :9300          ┃
┃ COMMANDS   │ gh workflow run deploy_ui.yml --ref feat/login-matrix                       ┃
┃            │ gh run watch $(gh run list --workflow deploy_ui.yml -q '.[0].databaseId')   ┃
┃            │ restart‑ozon: ssh root systemctl restart ozon-bridge                        ┃
┃ SUB‑AGENTS │ ui‑linter · a11y‑bot · visual‑diff · copy‑guard · data‑mocker               ┃
┃ NEXT STEPS │ 1) merge feat/login-matrix→main                                             ┃
┃            │ 2) /dashboard/ozon SSE panel                                                ┃
┃            │ 3) refresh‑token flow + error UI                                            ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
