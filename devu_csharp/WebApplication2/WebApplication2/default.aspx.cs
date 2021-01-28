using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace WebApplication2
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

            // initialize calendars if not page postback. 
            if (!Page.IsPostBack)
            {
                endPreviousAssignCalendar.SelectedDate = DateTime.Now.Date;
                startNewAssignCalendar.SelectedDate = DateTime.Now.Date.AddDays(14);
                endNewAssignCalendar.SelectedDate = DateTime.Now.Date.AddDays(21);
                               
            }
        }

        protected void Button1_Click(object sender, EventArgs e)
        {

           
            TimeSpan spanMission = startNewAssignCalendar.SelectedDate.Date.Subtract(endPreviousAssignCalendar.SelectedDate.Date);

            if (spanMission.TotalDays < 14)
            {
                resultLabel.Text = "Error: Must allow at least two weeks between previous assignment and new assignment";
                startNewAssignCalendar.SelectedDate = endPreviousAssignCalendar.SelectedDate.Date.AddDays(14);
            }
            else
            {
                TimeSpan projectMission = endNewAssignCalendar.SelectedDate.Date.Subtract(startNewAssignCalendar.SelectedDate.Date);

                double budget = projectMission.TotalDays * 500;

                budget = (projectMission.TotalDays > 21) ? budget += 1000 : budget;

                string result = String.Format("Assignment {0} to assignment project: {1} is authorized. Budget total: {2:C}",
                    spyCodeNameTextBox.Text,
                    newAssignmentTextBox.Text,
                    budget.ToString());

                // Project mission Total days is only to test the solution. 
                // In the final app you need to suppress it from resultLabel.Text
                resultLabel.Text = result + " " + projectMission.TotalDays.ToString();
            }

        }
    }
}