//StandaloneDashboard(function(db){
//
//	db.setDashboardTitle ("My Dashboard");
//
//	var chart = new ChartComponent();
//	chart.setCaption("Sales");
//	chart.setDimensions (6, 6);
//	chart.setLabels (["2013", "2014", "2015"]);
//	chart.addSeries ([3151, 1121, 4982]);
//	db.addComponent (chart);
//
//	var chart2 = new ChartComponent();
//	chart2.setCaption("Sales");
//	chart2.setDimensions (6, 6);
//	chart2.setLabels (["2013", "2014", "2015"]);
//	chart2.addSeries ([3151, 1121, 4982], {
//		numberPrefix: "$",
//		seriesDisplayType: "line"
//	});
//	db.addComponent (chart2);
//});
//

$.ajax({
    type: "GET",
    dataType: "json",
    url: "http://10.116.136.13:8003/line",
    success: function (data) {
        rf.StandaloneDashboard(function (db) {
        var sales_chart = new ChartComponent();
        sales_chart.setCaption("Sales Chart");
        db.addComponent(sales_chart);

        sales_chart.setLabels (data['Categories']);
        sales_chart.addSeries ("Sales", "sales", data['Sales']);

        });
    }
})


//
//rf.StandaloneDashboard(function(db){
//    var sales_chart = new ChartComponent();
//    sales_chart.setCaption ("Sales for 2014");
//    sales_chart.lock ();
//    db.addComponent(sales_chart);
//
//    $.get("http://localhost:8003/line", function (data) {
//        // This function is executed when the ajax request is successful.
//
//        sales_chart.setLabels (data['Categories']); // You can also use data.categories
//        sales_chart.addSeries ("Sales", "sales", data['Sales']);
//
//        // Don't forget to call unlock or the data won't be displayed
//        sales_chart.unlock ();
//    });
//});
//
