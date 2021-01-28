using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace pageLoad
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {
            if (!Page.IsPostBack)
            { 
            myTextBox.Text = "Some Value";
            myCalendar.SelectedDate = DateTime.Now.Date.AddDays(2);
            }
        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            resultLabel.Text = myTextBox.Text + " - "
                + myCalendar.SelectedDate.ToShortDateString();

        }
    }
}