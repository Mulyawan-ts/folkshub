import flet as ft
from components.navbar import Navbar


def HomePage(page: ft.Page):
    return ft.View(
        route="/",
        controls=[
            Navbar("FolksHub - Dashboard"),
            ft.Container(
                content=ft.Column(
                    controls=[
                        ft.Text(
                            "Selamat datang di FolksHub!",
                            size=24,
                            weight=ft.FontWeight.BOLD,
                        ),
                        ft.Text("Ini adalah tampilan halaman utama."),
                    ]
                ),
                padding=20,
            ),
        ],
    )
