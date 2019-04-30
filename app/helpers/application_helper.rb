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

  def decorated
    resource.decorate
  end

  def breadcrumbs
    content_tag :nav, class: :breadcrumbs do
      content_tag :div, class: 'nav-wrapper' do
        content_tag :div, class: 'col s12' do
          concat link_to('Backoffice', :backoffice, class: 'breadcrumb', data: { remote: true })

          yield if block_given?
        end
      end
    end
  end
end
