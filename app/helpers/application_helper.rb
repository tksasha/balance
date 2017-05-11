module ApplicationHelper
  def current_date
    @current_date ||= DateFactory.build params
  end

  def months
    %w(Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень)
  end

  def money sum
    number_with_delimiter '%.2f' % sum
  end

  #
  # TODO: deprecated
  #
  def cashes
    Cash.order(:name)
  end

  def decorated
    resource.decorate
  end
end
