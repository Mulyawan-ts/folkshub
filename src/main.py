import flet as ft
from assets.theme import setup_theme
from pages.home import HomePage
from src.components.button import PrimaryButton

def main(page: ft.Page):
    setup_theme(page)

def on_btn_click(e):
        page.add(ft.Text("Tombol berhasil diklik!"))

    page.add(
        ft.Container(
            content=ft.Column(
                controls=[
                    ft.Text("Uji Coba Komponen Button", size=20, weight=ft.FontWeight.BOLD),
                    PrimaryButton(
                        text="Klik Saya",
                        icon=ft.icons.CHECK,
                        on_click=on_btn_click
                    )
                ]
            ),
            padding=20
        )
    )

if __name__ == "__main__":
    ft.app(target=main)
