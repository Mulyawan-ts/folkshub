import flet as ft


def PrimaryButton(text: str, on_click=None, icon=None):
    return ft.ElevatedButton(
        text=text,
        icon=icon,
        on_click=on_click,
        style=ft.ButtonStyle(
            color=ft.colors.WHITE,
            bgcolor=ft.colors.BLUE_600,
            padding=ft.padding.all(15),
            shape=ft.RoundedRectangleBorder(radius=8),
        ),
    )
