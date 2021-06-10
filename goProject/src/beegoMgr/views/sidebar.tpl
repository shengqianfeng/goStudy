{{define "sidebar"}}
<!--左侧导航菜单-->
<div class="sidebar-nav">
    <ul>
    {{range .Menus}}
        <li>
            <a href="#" data-target=".dashboard-menu-{{.Id}}" class="nav-header" data-toggle="collapse">
            <i class="fa fa-fw fa-dashboard"></i> {{.MenuName}}<i class="fa fa-collapse"></i>
            </a>
        </li>
        <li>
            <ul class="dashboard-menu-{{.Id}} nav nav-list collapse in">
            {{range .SMenus}}
                <li><a href="{{.Url}}" target="mainFrame"><span class="fa fa-caret-right"></span>{{.MenuName}}</a></li>
            {{end}}
            </ul>
        </li>
    {{end}}
    </ul>
</div>
{{end}}