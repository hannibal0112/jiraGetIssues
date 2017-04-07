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
    c3.setSparkValues(['Jan', "Feb", 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],[12.31, 10.34, 10.26, 9, 8.21, 13.41, 14.43, 23.31, 13.41, 11.4, 28.34, 29.21]);
    db2.addComponent(c3);

    // Dashboard 3
    var db3 = new Dashboard('db3');
    db3.setDashboardTitle("Dashboard 3");

    var gauge = new GaugeComponent();
    gauge.setDimensions(4,3);
    gauge.setCaption('Gauge Component');
    gauge.setValue(144, {numberPrefix: '$'});
    gauge.setLimits(0, 200);
    db3.addComponent(gauge);

    var kpic = new KPIComponent ();
    kpic.setDimensions (4, 4);
    kpic.setCaption ("KPI Component");
    kpic.setValue (42);
    db3.addComponent(kpic);

    var kpi = new KPIGroupComponent ();
    kpi.setDimensions (12, 2);
    kpi.setCaption('KPI Group Component');
    kpi.addKPI('beverages', {
        caption: 'Beverages',
        value: 559,
        numberSuffix: ' units'
    });
    kpi.addKPI('condiments', {
        caption: 'Condiments',
        value: 507,
        numberSuffix: ' units'
    });
    kpi.addKPI('confections', {
        caption: 'Confections',
        value: 386,
        numberSuffix: ' units'
    });
    kpi.addKPI('daily_products', {
        caption: 'Daily Products',
        value: 393,
        numberSuffix: ' units'
    });
    db3.addComponent (kpi);

    var kpiT = new KPITableComponent ();
    kpiT.setDimensions (4, 6);
    kpiT.setCaption('KPI Table Component');
    kpiT.addKPI('grains_cereals', {
        caption: 'Grains/Cereals',
        value: 308,
        numberSuffix: ' units'
    });
    kpiT.addKPI('meat_poultry', {
        caption: 'Meat/Poultry',
        value: 165,
        numberSuffix: ' units'
    });
    kpiT.addKPI('produce', {
        caption: 'Produce',
        value: 100,
        numberSuffix: ' units'
    });
    kpiT.addKPI('seafood', {
        caption: 'Sea Food',
        value: 701,
        numberSuffix: ' units'
    });
    db3.addComponent (kpiT);

    var table = new TableComponent ('test');
    table.setCaption ("Table Component");
    table.setDimensions(4, 4);
    table.addColumn ('zone', "Zone");
    table.addColumn ('name', "Store Name");
    table.addColumn ('sale', "Sales amount");
    var data = [
        {zone: "North", name: "Northern Stores", sale: 4000},
        {zone: "South", name: "Southern Stores", sale: 4500},
    ];
    table.addMultipleRows (data);
    db3.addComponent(table);


    // Setup Dashboard Argument
    tdb.addDashboardTab(db1, {
        title: 'First Dashboard'
    });
    tdb.addDashboardTab(db2, {
        title: 'Second Dashboard'
    });
    tdb.addDashboardTab(db3, {
        title: 'Third Dashboard',
        active: true
    });

}, {tabbed: true});
