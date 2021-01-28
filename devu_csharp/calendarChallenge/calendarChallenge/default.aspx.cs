using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace calendarChallenge
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            // DateTime myDateOne = Calendar1.SelectedDate;
            // DateTime myDateTwo = Calendar2.SelectedDate;
            // TimeSpan result = myDateOne.Subtract(myDateTwo);

            if (Calendar1.SelectedDate > Calendar2.SelectedDate)
            {
                resultLabel.Text = Calendar1.SelectedDate
                    .Subtract(Calendar2.SelectedDate)
                    .TotalDays
                    .ToString();
                 
            }
            else
            {
                resultLabel.Text = Calendar2.SelectedDate
                   .Subtract(Calendar1.SelectedDate)
                   .TotalDays
                   .ToString();
            }



        }
    }
}