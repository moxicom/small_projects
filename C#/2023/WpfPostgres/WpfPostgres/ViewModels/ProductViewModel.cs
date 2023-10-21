using GalaSoft.MvvmLight;
using GalaSoft.MvvmLight.Command;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Linq;
using System.Reflection.Metadata;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;
using WpfPostgres.Data;
using WpfPostgres.Models;

namespace WpfPostgres.ViewModels
{
    public class ProductViewModel : ViewModelBase, INotifyPropertyChanged
    {
        private ObservableCollection<User> _users;
        private string _someString;

        public ICommand RemoveCommand { get; }

        public string SomeString
        {
            get => _someString;
            set
            {
                _someString = value;
                OnPropertyChanged(nameof(SomeString));
            }
        }

        public ObservableCollection<User> Users
        {
            get { return _users; }
            set
            {
                _users = value;
                OnPropertyChanged(nameof(Users));
            }
        }

        public ProductViewModel()
        {
            LoadData();
            RemoveCommand = new RelayCommand<int>(DeleteItem);
        }

        private void LoadData()
        {
            using (var dbContext = new ApplicationContext())
            {
                Users = new ObservableCollection<User>(dbContext.Users.ToList());
            }
        }

        public event PropertyChangedEventHandler? PropertyChanged;

        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }

        public void DeleteItem(int parameter)
        {
            SomeString = parameter.ToString();
        }
    }
}

