# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  before { travel_to Date.new 2019, 10, 14 }

  after { travel_back }

  describe '#search_by_date_range' do
    context do
      let :date_range do
        Date.new(2019, 10, 1)..Date.new(2019, 10, 31)
      end

      before { expect(relation).to receive(:where).with(date: date_range).and_return(:collection) }

      subject { described_class.new relation }

      its(:search) { should eq :collection }
    end

    context do
      let :date_range do
        Date.new(2019, 10, 1)..Date.new(2019, 10, 31)
      end

      before { expect(relation).to receive(:where).with(date: date_range).and_return(:collection) }

      subject { described_class.new relation, year: 2019 }

      its(:search) { should eq :collection }
    end

    context do
      let :date_range do
        Date.new(2018, 10, 1)..Date.new(2018, 10, 31)
      end

      before { expect(relation).to receive(:where).with(date: date_range).and_return(:collection) }

      subject { described_class.new relation, year: 2018 }

      its(:search) { should eq :collection }
    end

    context do
      let :date_range do
        Date.new(2019, 9, 1)..Date.new(2019, 9, 30)
      end

      before { expect(relation).to receive(:where).with(date: date_range).and_return(:collection) }

      subject { described_class.new relation, year: 2019, month: 9 }

      its(:search) { should eq :collection }
    end

    context do
      let :date_range do
        Date.new(2018, 9, 1)..Date.new(2018, 9, 30)
      end

      before { expect(relation).to receive(:where).with(date: date_range).and_return(:collection) }

      subject { described_class.new relation, year: 2018, month: 9 }

      its(:search) { should eq :collection }
    end
  end
end
