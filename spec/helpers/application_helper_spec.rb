# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  its :months do
    should eq %w[Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень]
  end

  describe '#current_date' do
    let(:date) { Date.new 2013, 6, 27 }

    before { allow(helper).to receive(:params).and_return(year: '2013', month: '06', day: '27') }

    before { expect(DateFactory).to receive(:build).with(year: '2013', month: '06', day: '27').and_return(date) }

    subject { helper.current_date }

    it { should eq date }
  end

  describe '#money' do
    subject { helper.money 400_500.2 }

    it { should eq '400 500.20' }
  end

  describe '#decorated' do
    before { expect(helper).to receive_message_chain(:resource, :decorate).and_return(:decorated) }

    subject { helper.decorated }

    it { should eq :decorated }
  end

  describe '#category_widget_data' do
    subject { helper }

    context do
      before { subject.instance_variable_set :@category_widget_data, :category_widget_data }

      its(:category_widget_data) { should eq :category_widget_data }
    end

    context do
      let(:params) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(CategoryWidgetDataSearcher).to receive(:search).with(params).and_return(:category_widget_data) }

      its(:category_widget_data) { should eq :category_widget_data }
    end
  end

  describe '#at_end' do
    context do
      before { subject.instance_variable_set :@at_end, :at_end }

      its(:at_end) { should eq :at_end }
    end

    context do
      let(:params) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(AtEndCalculatorService).to receive(:calculate).with(params).and_return(21.49) }

      its(:at_end) { should eq 21.49 }
    end
  end

  describe '#balance' do
    context do
      before { subject.instance_variable_set :@balance, :balance }

      its(:balance) { should eq :balance }
    end

    context do
      let(:params) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(BalanceCalculatorService).to receive(:calculate).with(params).and_return(22.15) }

      its(:balance) { should eq 22.15 }
    end
  end

  describe '#new_item_for_inline_form' do
    let(:params) { { currency: 'usd', foo: 'foo' } }

    before { expect(subject).to receive(:params).and_return(params) }

    before { expect(Item).to receive(:new).with(currency: 'usd').and_return(:item) }

    its(:new_item_for_inline_form) { should eq :item }
  end
end
