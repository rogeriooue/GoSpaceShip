# GoSpaceShip

2D Game in Go (Golang).

## Descrição

**GoSpaceShip** é um jogo 2D desenvolvido em Go utilizando a biblioteca [Ebiten](https://ebiten.org/). O objetivo do jogo é controlar uma nave espacial, desviar de meteoros e acumular pontos ao destruir os obstáculos.

---

## Funcionalidades

- Controle de uma nave espacial.
- Geração aleatória de meteoros.
- Sistema de pontuação.
- Tela de "Game Over" com opção de reiniciar o jogo.
- Gráficos 2D simples e eficientes.

---

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang).
- **Biblioteca de Jogos**: [Ebiten](https://ebiten.org/).
- **Gerenciamento de Dependências**: Go Modules.

---

## Como Jogar

1. **Iniciar o jogo**:
   - Execute o comando abaixo no terminal para iniciar o jogo:

     go run main.go

   - Ou através do executável GoSpaceShip.exe.

2. **Controles**:
   - **Setas direcionais**: Movem a nave espacial.
   - **Barra de espaço**: Dispara lasers para destruir meteoros.

3. **Objetivo**:
   - Desvie dos meteoros e acumule pontos ao destruí-los.

4. **Game Over**:
   - O jogo termina quando a nave colide com um meteoro.
   - Após 2 segundos, pressione a barra de espaço para reiniciar o jogo.
