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

  def breadcrumbs
    content_tag :nav, class: :breadcrumb do
      concat link_to('Backoffice', :backoffice, class: 'breadcrumb-item', data: { remote: true })

      yield if block_given?
    end
  end
end
