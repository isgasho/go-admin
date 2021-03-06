// Code generated by hero.
// source: /Users/chenhg5/go/src/goAdmin/resources/pages/new_render.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"goAdmin/menu"
	"goAdmin/models"

	"github.com/shiyanhui/hero"
)

func NewPanel(formData []models.FormStruct, url string, previous string, id string, menuList []menu.MenuItem, title string, description string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>GoAdmin</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="../../assets/bootstrap/dist/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="../../assets/font-awesome/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="../../assets/Ionicons/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="../../assets/datatables.net-bs/css/dataTables.bootstrap.min.css">
    <!-- iCheck -->
    <link rel="stylesheet" href="../../assets/iCheck/minimal/_all.css">
    <link rel="stylesheet" href="../../assets/iCheck/futurico/futurico.css">
    <link rel="stylesheet" href="../../assets/iCheck/polaris/polaris.css">
    <link rel="stylesheet" href="../../assets/toastr/build/toastr.min.css">
    <link rel="stylesheet" href="../../assets/nprogress/nprogress.css">
    <link rel="stylesheet" href="../../assets/select2/select2.min.css">
    <link rel="stylesheet" href="../../assets/fileinput/fileinput.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="../../assets/dist/css/AdminLTE.min.css">
    <!-- AdminLTE Skins. Choose a skin from the css/skins
         folder instead of downloading all of them to reduce the load. -->
    <link rel="stylesheet" href="../../assets/dist/css/skins/_all-skins.min.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="../../assets/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="../../assets/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link rel="stylesheet" href="../../assets/googleapis/font.css">
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">

    <header class="main-header">
        <!-- Logo -->
        <a href="../../index2.html" class="logo">
            <!-- mini logo for sidebar mini 50x50 pixels -->
            <span class="logo-mini"><b>G</b>A</span>
            <!-- logo for regular state and mobile devices -->
            <span class="logo-lg"><b>Go</b>Admin</span>
        </a>
        <!-- Header Navbar: style can be found in header.less -->
        <nav class="navbar navbar-static-top">
            <!-- Sidebar toggle button-->
            <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </a>

            <div class="navbar-custom-menu">
                <ul class="nav navbar-nav">
                    <!-- User Account: style can be found in dropdown.less -->
                    <li class="dropdown user user-menu">
                    <li class="dropdown user user-menu">
                        `)
	buffer.WriteString(`
                    </li>
                </ul>
            </div>
        </nav>
    </header>
    <!-- Left side column. contains the logo and sidebar -->
    <aside class="main-sidebar">
        <!-- sidebar: style can be found in sidebar.less -->
        <section class="sidebar">
            <!-- Sidebar user panel -->
            <div class="user-panel">
                <div class="pull-left image">
                    <img src="../../assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                </div>
                <div class="pull-left info">
                    `)
	buffer.WriteString(`
                    <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
                </div>
            </div>
            <!-- search form -->
            <form action="#" method="get" class="sidebar-form">
                <div class="input-group">
                    <input type="text" name="q" class="form-control" placeholder="Search...">
                    <span class="input-group-btn">
                <button type="submit" name="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
                </div>
            </form>
            <!-- /.search form -->
            <!-- sidebar menu: : style can be found in sidebar.less -->
            <ul class="sidebar-menu" data-widget="tree">
                <li class="header">MAIN NAVIGATION</li>
                <!--<li class="treeview">
                  <a href="#">
                    <i class="fa fa-edit"></i> <span>Forms</span>
                    <span class="pull-right-container">
                          <i class="fa fa-angle-left pull-right"></i>
                        </span>
                  </a>
                  <ul class="treeview-menu">
                    <li><a href="../forms/general.html"><i class="fa fa-circle-o"></i> General Elements</a></li>
                    <li><a href="../forms/advanced.html"><i class="fa fa-circle-o"></i> Advanced Elements</a></li>
                    <li><a href="../forms/editors.html"><i class="fa fa-circle-o"></i> Editors</a></li>
                  </ul>
                </li>
                <li class="treeview active">
                  <a href="#">
                    <i class="fa fa-table"></i> <span>Tables</span>
                    <span class="pull-right-container">
                          <i class="fa fa-angle-left pull-right"></i>
                        </span>
                  </a>
                  <ul class="treeview-menu">
                    <li><a href="simple.html"><i class="fa fa-circle-o"></i> Simple tables</a></li>
                    <li class="active"><a href="data.html"><i class="fa fa-circle-o"></i> Data tables</a></li>
                  </ul>
                </li>-->
                `)
	for _, list := range menuList {
		if len(list.ChildrenList) == 0 {
			buffer.WriteString(`
<li class='`)
			hero.EscapeHTML(list.Active, buffer)
			buffer.WriteString(`'>
    <a href='`)
			hero.EscapeHTML(list.Url, buffer)
			buffer.WriteString(`'>
        <i class="fa `)
			hero.EscapeHTML(list.Icon, buffer)
			buffer.WriteString(`"></i><span>`)
			hero.EscapeHTML(list.Name, buffer)
			buffer.WriteString(`</span>
        <span class="pull-right-container"><!-- <small class="label pull-right bg-green">new</small> --></span>
    </a>
</li>
`)
		} else {
			buffer.WriteString(`
<li class="treeview `)
			hero.EscapeHTML(list.Active, buffer)
			buffer.WriteString(`">
    <a href="#">
        <i class="fa `)
			hero.EscapeHTML(list.Icon, buffer)
			buffer.WriteString(`"></i> <span>`)
			hero.EscapeHTML(list.Name, buffer)
			buffer.WriteString(`</span>
        <span class="pull-right-container">
                                  <i class="fa fa-angle-left pull-right"></i>
                                </span>
    </a>
    <ul class="treeview-menu">
        `)
			for _, item := range list.ChildrenList {
				buffer.WriteString(`
        <li><a href="`)
				hero.EscapeHTML(item.Url, buffer)
				buffer.WriteString(`"><i class="fa `)
				hero.EscapeHTML(item.Icon, buffer)
				buffer.WriteString(`"></i> `)
				hero.EscapeHTML(item.Name, buffer)
				buffer.WriteString(`</a></li>
        `)
			}
			buffer.WriteString(`
    </ul>
</li>
`)
		}
	}

	buffer.WriteString(`
            </ul>
        </section>
        <!-- /.sidebar -->
    </aside>

    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper" id="pjax-container">
        <section class="content-header">
            `)
	buffer.WriteString(`
<h1>
    `)
	hero.EscapeHTML(title, buffer)
	buffer.WriteString(`
    <small>`)
	hero.EscapeHTML(description, buffer)
	buffer.WriteString(`</small>
</h1>
`)

	buffer.WriteString(`
            <!-- breadcrumb start -->
            <!-- breadcrumb end -->
        </section>
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box box-info">
                        <div class="box-header with-border">
                            <h3 class="box-title">Edit</h3>
                            <div class="box-tools">
                                <div class="btn-group pull-right" style="margin-right: 10px">
                                    <a href="" class="btn btn-sm btn-default"><i class="fa fa-list"></i> List</a>
                                </div>
                                <div class="btn-group pull-right" style="margin-right: 10px">
                                    <a class="btn btn-sm btn-default form-history-back"><i class="fa fa-arrow-left"></i> Back</a>
                                </div>
                            </div>
                        </div>
                        <!-- /.box-header -->
                        <!-- form start -->
                        `)
	buffer.WriteString(`
<form action='`)
	hero.EscapeHTML(url, buffer)
	buffer.WriteString(`' method="post" accept-charset="UTF-8" class="form-horizontal" pjax-container>
`)

	buffer.WriteString(`
                        <div class="box-body">
                            <div class="fields-group">
                                `)
	for _, data := range formData {
		buffer.WriteString(`
<div class="form-group ">
    `)
		if data.Field != "created_at" && data.Field != "updated_at" && data.Field != "id" {
			if data.FormType == "default" {
				buffer.WriteString(`
        <label class="col-sm-2 control-label">`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`</label>
        <div class="col-sm-8">
            <div class="box box-solid box-default no-margin">
                <!-- /.box-header -->
                <div class="box-body">
                    `)
				hero.EscapeHTML(data.Value, buffer)
				buffer.WriteString(` 
                </div>
                <!-- /.box-body -->
            </div>
        </div>
    `)
			} else if data.FormType == "text" {
				buffer.WriteString(`
        <label for="json" class="col-sm-2 control-label">`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`</label>
        <div class="col-sm-8">
            <div class="input-group">
                <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
                <input type="text" id="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`" name="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`" value='' class="form-control json" placeholder="Input json">
            </div>
        </div>
    `)
			} else if data.FormType == "password" {
				buffer.WriteString(`
        <label for="password" class="col-sm-2 control-label">`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`</label>
        <div class="col-sm-8">
            <div class="input-group">
                <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
                <input type="password" id="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`" name="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`" value="" class="form-control password" placeholder="Input Password">
            </div>
        </div>
    `)
			} else if data.FormType == "textarea" {
				buffer.WriteString(`
        <label for="http_path" class="col-sm-2 control-label">`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`</label>
        <div class="col-sm-8">
            <textarea name="http_path" class="form-control" rows="5" placeholder="Input text"></textarea>
        </div>
    `)
			} else if data.FormType == "select" {
				buffer.WriteString(`
        <label for="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`" class="col-sm-2 control-label">`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`</label>
        <div class="col-sm-8">
            <select class="form-control http_method select2-hidden-accessible" style="width: 100%;" name="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`[]" multiple="" data-placeholder="Input HTTP method" tabindex="-1" aria-hidden="true">
                `)
				for _, v := range data.Options {
					buffer.WriteString(`
                    <option value='`)
					hero.EscapeHTML(v["value"], buffer)
					buffer.WriteString(`'>`)
					hero.EscapeHTML(v["field"], buffer)
					buffer.WriteString(`</option>
                `)
				}
				buffer.WriteString(`
            </select>
            <input type="hidden" name="`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`[]">
            <span class="help-block">
                    <i class="fa fa-info-circle"></i>&nbsp;All methods if empty
                </span>
        </div>
        <script>
            $(".`)
				hero.EscapeHTML(data.Field, buffer)
				buffer.WriteString(`").select2({
                allowClear: true,
                placeholder: "`)
				hero.EscapeHTML(data.Head, buffer)
				buffer.WriteString(`"
            });
        </script>
    `)
			}
		}
		buffer.WriteString(`
</div>
`)
	}

	buffer.WriteString(`

                                <!-- <div class="form-group ">
                                    <label class="col-sm-2 control-label">ID</label>
                                    <div class="col-sm-8">
                                        <div class="box box-solid box-default no-margin">
                                            <div class="box-body">
                                                1 
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="form-group  ">
                                    <label for="json" class="col-sm-2 control-label">json</label>
                                    <div class="col-sm-8">
                                        <div class="input-group">
                                            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>

                                            <input type="text" id="json" name="json" value='123' class="form-control json" placeholder="Input json">
                                        </div>
                                    </div>
                                </div>
                -->

                            </div>

                        </div>
                        <!-- /.box-body -->
                        <div class="box-footer">
                            <input type="hidden" name="_token" value="7TEJrUaKsAIZ0qbc03G1nmeDOfmHyhCbMHqHlnkg">
                            <div class="col-md-2">
                            </div>
                            <div class="col-md-8">

                                <div class="btn-group pull-right">
                                    <button type="submit" class="btn btn-info pull-right"
                                            data-loading-text="&lt;i class='fa fa-spinner fa-spin '&gt;&lt;/i&gt; Save">
                                        Save
                                    </button>
                                </div>

                                <div class="btn-group pull-left">
                                    <button type="reset" class="btn btn-warning">Reset</button>
                                </div>

                            </div>

                        </div>

                        <input type="hidden" name="_method" value="PUT" class="_method">
                        `)
	buffer.WriteString(`
<input type="hidden" name="_previous_" value='`)
	hero.EscapeHTML(previous, buffer)
	buffer.WriteString(`' class="_previous_">
<input type="hidden" name="id" value='`)
	hero.EscapeHTML(id, buffer)
	buffer.WriteString(`' class="_previous_">
`)

	buffer.WriteString(`
                        <!-- /.box-footer -->
                        </form>
                    </div>

                </div>
            </div>

        </section>
        <script data-exec-on-popstate>

            $(function () {
                $('.form-history-back').on('click', function (event) {
                    event.preventDefault();
                    history.back(1);
                });
            });
        </script>
    </div>
    <!-- /.content-wrapper -->
    <footer class="main-footer">
        <div class="pull-right hidden-xs">
            <b>Version</b> 2.4.0
        </div>
        <strong>Copyright &copy; 2018- <a href="https://github.com/chenhg5/go-admin">GoAdmin</a>.</strong> All rights
        reserved.
    </footer>
</div>
<!-- ./wrapper -->

<!-- jQuery 3 -->
<!-- <script src="../../assets/jquery/dist/jquery.min.js"></script> -->
<script src="../../assets/jQuery/jQuery-2.1.4.min.js"></script>
<!-- Bootstrap 3.3.7 -->
<script src="../../assets/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="../../assets/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="../../assets/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="../../assets/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="../../assets/fastclick/lib/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="../../assets/dist/js/adminlte.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="../../assets/dist/js/demo.js"></script>
<script src="../../assets/select2/select2.full.min.js"></script>
<script src="../../assets/fileinput/fileinput.min.js"></script>
<script src="../../assets/iCheck/icheck.min.js"></script>
<script src="../../assets/nprogress/nprogress.js"></script>
<script src="../../assets/toastr/build/toastr.min.js"></script>
<script src="../../assets/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
<script src="../../assets/jquery-pjax/jquery.pjax.js"></script>
<script>
    $('.grid-per-pager').on("change", function (e) {
        console.log("changing...")
        $.pjax({url: this.value, container: '#pjax-container'});
    });
    $('.grid-refresh').on('click', function () {
        $.pjax.reload('#pjax-container');
        toastr.success('Refresh succeeded !');
    });
    $.fn.editable.defaults.params = function (params) {
        params._token = LA.token;
        params._editable = 1;
        params._method = 'PUT';
        return params;
    };

    $.fn.editable.defaults.error = function (data) {
        var msg = '';
        if (data.responseJSON.errors) {
            $.each(data.responseJSON.errors, function (k, v) {
                msg += v + "\n";
            });
        }
        return msg
    };

    toastr.options = {
        closeButton: true,
        progressBar: true,
        showMethod: 'slideDown',
        timeOut: 4000
    };

    $.pjax.defaults.timeout = 5000;
    $.pjax.defaults.maxCacheLength = 0;
    $(document).pjax('a:not(a[target="_blank"])', {
        container: '#pjax-container'
    });

    NProgress.configure({parent: '#pjax-container'});

    $(document).on('pjax:timeout', function (event) {
        event.preventDefault();
    })

    $(document).on('submit', 'form[pjax-container]', function (event) {
        $.pjax.submit(event, '#pjax-container')
    });

    $(document).on("pjax:popstate", function () {

        $(document).one("pjax:end", function (event) {
            $(event.target).find("script[data-exec-on-popstate]").each(function () {
                $.globalEval(this.text || this.textContent || this.innerHTML || '');
            });
        });
    });

    $(document).on('pjax:send', function (xhr) {
        if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
            $submit_btn = $('form[pjax-container] :submit');
            if ($submit_btn) {
                $submit_btn.button('loading')
            }
        }
        NProgress.start();
    });

    $(document).on('pjax:complete', function (xhr) {
        if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
            $submit_btn = $('form[pjax-container] :submit');
            if ($submit_btn) {
                $submit_btn.button('reset')
            }
        }
        NProgress.done();
    });

    $(function () {
        $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
            var $parent = $(this).parent().addClass('active');
            $parent.siblings('.treeview.active').find('> a').trigger('click');
            $parent.siblings().removeClass('active').find('li').removeClass('active');
        });

        $('[data-toggle="popover"]').popover();
    });

    // (function ($) {
    //     $.fn.admin = LA;
    //     $.admin = LA;
    //
    // })(jQuery);
</script>
</body>
</html>
`)

}
