StandaloneDashboard(function (tdb) {
    tdb.setTabbedDashboardTitle("Tabbed Dashboard");

    // Dashboard 1
    var db1 = new Dashboard();
    db1.setDashboardTitle('Table In Razorfow');

    var sales_chart = new ChartComponent();
    sales_chart.setCaption("Sales Chart");

    sales_chart.lock();
    db1.addComponent(sales_chart);
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "http://10.116.136.13:8003/line/TYGH",
        success: function (data) {
            sales_chart.setLabels (data['Categories']);
            sales_chart.addSeries ("Sales", "sales", data['Sales']);
            sales_chart.unlock();
        }
    });

    // Dashboard 2
    var db2 = new Dashboard('db2');
    db2.setDashboardTitle('KPI Types Supported in RazorFlow');

    var c2 = new KPIComponent();
    c2.setDimensions(4, 2);
    c2.setCaption('Average Monthly Sales');
    c2.setValue(513.22, {
        numberPrefix: '$'
    });
    db2.addComponent(c2);

    var c3 = new KPIComponent();
    c3.setDimensions(4, 2);
    c3.setCaption('Average Monthly Units');
    c3.setValue(22);
    c3.setSparkValues(['Jan', "Feb", 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
                      [12.31, 10.34, 10.26, 9, 8.21, 13.41, 14.43, 23.31, 13.41, 11.4, 28.34, 29.21]);
    db2.addComponent(c3);

    tdb.addDashboardTab(db1, {
        title: 'First Dashboard'
    });
    tdb.addDashboardTab(db2, {
        title: 'Second Dashboard',
        active: true
    });

}, {tabbed: true});
