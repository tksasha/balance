require 'rails_helper'

RSpec.describe ApplicationHelper, type: :helper do
  describe '#months' do
    subject { helper.months }

    it { should eq %w(Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень) }
  end

  describe '#current_date' do
    let(:date) { Date.new 2013, 6, 27 }

    before { allow(helper).to receive(:params) { { year: '2013', month: '06', day: '27' } } }

    before { expect(DateFactory).to receive(:build).with(year: '2013', month: '06', day: '27') { date } }

    subject { helper.current_date }

    it { should eq date }
  end

  describe '#money' do
    subject { helper.money 400_500.2 }

    it { should eq '400 500.20' }
  end

  describe '#decorated' do
    before do
      #
      # helper.resource.decorate -> :decorated
      #
      expect(helper).to receive(:resource) do
        double.tap { |a| expect(a).to receive(:decorate).and_return(:decorated) }
      end
    end

    subject { helper.decorated }

    it { should eq :decorated }
  end
end
