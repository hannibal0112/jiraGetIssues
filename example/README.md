# Usage for this example 

## File Stucture 
* rd ==> RazorFlow Server by JavaScript (link https://razorflow.com/docs/dashboard/js/guide/index.php)
* restServer.go ==>  a simple rest server for testing 
  
## rf/js/dashboard_app.js 
Get data of chart from your REST server. 
``` javascript
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
## restServer.go 
Easy REST server 
## JSON Structure 
```
{
  "Categories":["One","Two","Three","Foru","Five","Six","Seven","Eight"],
  "Sales":[1,2,3,4,5,6,7,8]
}
```
