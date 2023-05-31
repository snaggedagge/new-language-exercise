
mod game {

    pub struct TicTacToe {
        grid: usize,
        private_data: TicTacToeData
    }

    struct TicTacToeData {
        playing_field: Vec<Vec<char>>
    }

    impl TicTacToe {
        pub(crate) fn new(grid: usize) -> Self {
            let mut field = vec![];
            for x in 0..grid {
                field.push(vec![]);
                for _y in 0..grid {
                    field[x].push('X');
                }
            }
            Self {
                grid,
                private_data: TicTacToeData {playing_field: field},
            }
        }

        pub fn print_playing_field(&self) {
            for _x in 0..self.grid {
                print!("----");
            }
            println!();
            for x in 0..self.grid {
                for y in 0..self.grid {
                    print!(" {} {}", self.private_data.playing_field[x][y], if y < self.grid-1 { "|" } else { "" });
                }
                println!();
                if x < self.grid-1 {
                    println!();
                }
            }
            for _x in 0..self.grid {
                print!("----");
            }
            println!();
        }

        pub fn player_move(&mut self, player: char, tuple: usize, row:usize) {
            self.private_data.playing_field[self.grid - row][tuple] = player;
        }
    }
}

use game::TicTacToe;

fn main() {
    // Not completed, this is just a simple start to get a feeling for rust.
    let mut game = TicTacToe::new(5);
    game.print_playing_field();
    game.player_move('O', 2, 1);
    game.player_move('O', 2, 3);
    game.print_playing_field();
}