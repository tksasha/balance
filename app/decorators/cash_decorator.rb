class CashDecorator < Draper::Decorator
  delegate_all

  def css_class
    favorite? ? 'cash cash-favorite' : 'cash'
  end
end
