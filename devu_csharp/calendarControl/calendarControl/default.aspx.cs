using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace calendarControl
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {
            myCalendar.SelectedDate = DateTime.Now;
        }

        protected void getDateButton_Click(object sender, EventArgs e)
        {
            resultLabel.Text = myCalendar.SelectedDate.ToShortDateString();

        }

        protected void setDateButton_Click(object sender, EventArgs e)
        {
           myCalendar.SelectedDate = DateTime.Parse(setDateTextBox.Text);

          
        }

        protected void showDateButton_Click(object sender, EventArgs e)
        {
           myCalendar.VisibleDate = DateTime.Parse(setDateTextBox.Text);

            

        }

        protected void myCalendar_SelectionChanged(object sender, EventArgs e)
        {
            resultLabel.Text = myCalendar.SelectedDate.ToShortDateString();
        }
    }
}