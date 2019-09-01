# frozen_string_literal: true

RSpec.describe DateFactory do
  describe '#build' do
    before { allow(Date).to receive(:today) { Date.new(2013, 3, 31) } }

    its(:build) { should eq Date.new 2013, 3, 1 }

    context do
      subject { described_class.new year: '2011' }

      its(:build) { should eq Date.new(2011, 3, 1) }
    end

    context do
      subject { described_class.new month: '12' }

      its(:build) { should eq Date.new(2013, 12, 1) }
    end

    context do
      subject { described_class.new day: '17' }

      its(:build) { should eq Date.new(2013, 3, 17) }
    end

    context do
      subject { described_class.new year: 2013, month: 2 }

      its(:build) { should eq Date.new(2013, 2, 1) }
    end

    context do
      subject { described_class.new year: 2013, month: 2, day: 31 }

      its(:build) { should eq Date.new(2013, 3, 31) }
    end
  end

  describe '.build' do
    after { described_class.build(:params) }

    it do
      #
      # described_class.new(:params).build
      #
      expect(described_class).to receive(:new).with(:params) do
        double.tap { |a| expect(a).to receive(:build) }
      end
    end
  end
end
