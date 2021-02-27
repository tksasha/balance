# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#category_widget_data' do
    subject { helper }

    context do
      before { subject.instance_variable_set :@category_widget_data, :category_widget_data }

      its(:category_widget_data) { should eq :category_widget_data }
    end

    context do
      let(:params) { double }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(CategoryWidgetDataSearcher).to receive(:search).with(params).and_return(:category_widget_data) }

      its(:category_widget_data) { should eq :category_widget_data }
    end
  end
end
