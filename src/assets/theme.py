import flet as ft


def setup_theme(page: ft.Page):
    page.title = "Folkshub Desktop"
    page.theme_mode = ft.ThemeMode.DARK
    page.padding = 0
    page.window.width = 800
    page.window.height = 600
    page.bgcolor = ft.colors.GREY_900

    # Kustomisasi skema warna aplikasi
    page.theme = ft.Theme(
        color_scheme=ft.ColorScheme(
            primary=ft.colors.BLUE_400,
            surface=ft.colors.GREY_800,
        )
    )
