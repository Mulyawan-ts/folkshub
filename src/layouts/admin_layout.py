import flet as ft


def AdminLayout(page: ft.Page, body: ft.Control, selected_index: int = 0):
    def on_nav_change(e):
        idx = e.control.selected_index
        if idx == 0:
            page.go("/admin/dashboard")
        elif idx == 1:
            page.go("/admin/users")
        elif idx == 2:
            page.go("/admin/reports")
        elif idx == 3:
            page.go("/admin/settings")

    rail = ft.NavigationRail(
        selected_index=selected_index,
        label_type=ft.NavigationRailLabelType.ALL,
        min_width=100,
        destinations=[
            ft.NavigationRailDestination(
                icon=ft.icons.DASHBOARD_OUTLINED,
                selected_icon=ft.icons.DASHBOARD,
                label="Dashboard",
            ),
            ft.NavigationRailDestination(
                icon=ft.icons.PEOPLE_OUTLINED,
                selected_icon=ft.icons.PEOPLE,
                label="Pengguna",
            ),
            ft.NavigationRailDestination(
                icon=ft.icons.ASSESSMENT_OUTLINED,
                selected_icon=ft.icons.ASSESSMENT,
                label="Laporan",
            ),
            ft.NavigationRailDestination(
                icon=ft.icons.SETTINGS_OUTLINED,
                selected_icon=ft.icons.SETTINGS,
                label="Pengaturan",
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
