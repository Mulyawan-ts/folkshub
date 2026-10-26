import flet as ft


def UserLayout(page: ft.Page, body: ft.Control, selected_index: int = 0):
    def on_nav_change(e):
        idx = e.control.selected_index
        if idx == 0:
            page.go("/user/dashboard")
        elif idx == 1:
            page.go("/user/tasks")
        elif idx == 2:
            page.go("/user/profile")

    rail = ft.NavigationRail(
        selected_index=selected_index,
        label_type=ft.NavigationRailLabelType.ALL,
        min_width=100,
        destinations=[
            ft.NavigationRailDestination(
                icon=ft.icons.HOME_OUTLINED,
                selected_icon=ft.icons.HOME,
                label="Beranda",
            ),
            ft.NavigationRailDestination(
                icon=ft.icons.TASK_OUTLINED,
                selected_icon=ft.icons.TASK,
                label="Tugas",
            ),
            ft.NavigationRailDestination(
                icon=ft.icons.PERSON_OUTLINED,
                selected_icon=ft.icons.PERSON,
                label="Profil",
            ),
        ],
        on_change=on_nav_change,
    )

    return ft.Row(
        controls=[
            rail,
            ft.VerticalDivider(width=1),
            ft.Container(content=body, expand=True),
        ],
        expand=True,
    )
