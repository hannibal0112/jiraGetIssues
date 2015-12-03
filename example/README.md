# Usage for this example 

## File Stucture 
* rd ==> RazorFlow Server by JavaScript 
* restServer.go ==>  a simple rest server for testing 
  
#<rf/js/dashboard_app.js> 

```
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
```
