# frozen_string_literal: true

RSpec.describe ConsolidationDecorator do
  let(:consolidation) { stub_model Consolidation }

  let(:params) { {} }

  subject { consolidation.decorate context: params }

  it { should delegate_method(:year).to(:date) }

  it { should delegate_method(:month).to(:date) }

  describe '#date' do
    before { travel_to Date.new(2019, 11, 5) }

    after { travel_back }

    its(:date) { should eq Date.new(2019, 11, 1) }

    context do
      let(:params) { { year: 2018 } }

      its(:date) { should eq Date.new(2018, 11, 1) }
    end

    context do
      let(:params) { { month: 10 } }

      its(:date) { should eq Date.new(2019, 10, 1) }
    end

    context do
      let(:params) { { year: 2018, month: 10 } }

      its(:date) { should eq Date.new(2018, 10, 1) }
    end

    context do
      before { subject.instance_variable_set :@date, Date.new(2019, 1, 1) }

      its(:date) { should eq Date.new(2019, 1, 1) }
    end
  end
end
